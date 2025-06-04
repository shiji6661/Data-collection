package dao_mysql

import (
	"Data-collection/way/groupBuyingProduct"
	"common/global"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"models/model_product/model_mysql"
	"weikang/Data-collection/srv/product_srv/proto_product/product"

	"time"
)

const (
	IsTpl    = 0 // 是否发送模版消息："拼团成功！" 0未发送 1已发送
	IsRefund = 0 // 是否退款 0 未退款 1已退款
	Status   = 1 //拼团的状态 1进行中 2已完成 3未完成
	OK       = 2 // 已完成
	Fail     = 3 // 未完成
)

// todo:查询用户是否已经参加过该拼团
func FindGroupBuyingByUidAndCid(uid, cid int64) (gb *model_mysql.GroupBuying, err error) {
	gb = &model_mysql.GroupBuying{}
	err = gb.FindGroupBuyingByUidAndCid(uid, cid)
	if err != nil {
		return nil, err
	}
	return gb, nil
}

// todo:拼团完成之后，更新拼团状态
func UpdateGroupBuyingStatus(cid int64, status int) error {
	gb := &model_mysql.GroupBuying{}
	gb.UpdateGroupBuyingStatus(cid, status)
	return nil
}

// GenerateGroupUUID 生成拼团唯一标识
func GenerateGroupUUID() string {
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		// 处理错误，例如返回一个默认值
		return "default_uuid"
	}
	return uuidObj.String()
}

// 用户发起拼团
func CreateGroupBuying(in *product.CreateUserGroupRequest) (gb *model_mysql.GroupBuying, err error) {
	// 查找拼团商品信息
	id, err := FindGProductById(in.Cid)
	if err != nil {
		return nil, err
	}
	if id.ID == 0 {
		zap.L().Info("拼团商品不存在！请重新选择！")
		return nil, errors.New("拼团商品不存在！请重新选择！")
	}

	// 检查是否存在相同商品的拼团记录
	var existingGroup model_mysql.GroupBuying
	err = global.DB.Where("uid =? AND cid =?", in.Uid, in.Cid).First(&existingGroup).Error
	if err == nil {
		// 如果存在且状态允许重新发起拼团，可以考虑在这里添加逻辑
		// 目前简单处理为允许重新发起
		zap.L().Info("用户可以重新发起该商品的拼团")
	} else if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// 判断购买数量是否超出拼团商品库存
	if id.Stock <= 0 {
		message := fmt.Sprintf("%s 拼团活动库存不足！请重新选择！", id.Title)
		zap.L().Info(message)
		return nil, errors.New(message)
	}
	if id.Stock < in.Num {
		message := fmt.Sprintf("%s 拼团活动库存不足！请重新选择！", id.Title)
		zap.L().Info(message)
		return nil, errors.New(message)
	}

	// 判断时间是否过期
	now := time.Now().Format("2006-01-02 15:04:05")
	if now < id.StartTime {
		zap.L().Info("拼团活动未开始！请重新选择！")
		return nil, errors.New("拼团活动未开始！请重新选择！")
	}
	if now > id.StopTime {
		zap.L().Info("拼团活动已过期！请重新选择！")
		return nil, errors.New("拼团活动已过期！请重新选择！")
	}

	totalPrice := float64(in.Num) * id.Price
	gb = &model_mysql.GroupBuying{
		Uid:        in.Uid,
		TotalNum:   in.Num,
		TotalPrice: totalPrice,
		Cid:        in.Cid,
		Pid:        id.ProductId,
		People:     0,
		Price:      id.Price,
		StopTime:   id.StopTime,
		KId:        0,
		IsTpl:      IsTpl,
		IsRefund:   IsRefund,
		Status:     Status,
	}

	// 使用事务确保原子性
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 创建拼团记录
		if err := tx.Create(gb).Error; err != nil {
			return err
		}

		// 新拼团（团长），设置people为1，KId为0
		updateResult := tx.Model(&model_mysql.GroupBuying{}).
			Where("id =?", gb.Id).
			Updates(map[string]interface{}{
				"people": 1,
				"k_id":   0,
			})

		if updateResult.Error != nil {
			return updateResult.Error
		}

		if updateResult.RowsAffected == 0 {
			return errors.New("更新团长拼团记录people和k_id字段失败")
		}

		fmt.Println("新拼团创建，people设为1，k_id设为0")

		// 扣减库存
		_, err := groupBuyingProduct.ReduceGroupProduct(id.ProductId, in.Num)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, errors.New("创建拼团失败")
	}

	// 重新查询最新的拼团信息
	updatedGroup, err := FindUserGroupByCid(in.Cid)
	if err != nil {
		return nil, err
	}

	if updatedGroup.People < id.People {
		zap.L().Info("当前拼团人数不足！可以分享给更多好友！")
	} else {
		UpdateGroupBuyingStatus(in.Cid, OK)
		// 拼团已满，可添加拼团成功逻辑
		zap.L().Info("拼团已满！恭喜成功组团！")
	}

	return gb, nil
}
