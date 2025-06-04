package dao_mysql

import (
	"common/global"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"models/model_product/model_mysql"
	"product_srv/proto_product/product"
	"way/groupBuyingProduct"
)

// todo 用户加入拼团
func JoinGroupBuying(in *product.UserJoinGroupRequest, invitationCode string) (gb *model_mysql.GroupBuying, err error) {
	// 查找拼团商品信息
	id, err := FindGProductById(in.Cid)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	if id.ID == 0 {
		zap.L().Info("拼团商品不存在！请重新选择！")
		return nil, errors.New("拼团商品不存在！请重新选择！")
	}

	// 判断用户是否已经参加过该拼团
	cid, err := FindGroupBuyingByUidAndCid(in.Uid, in.Cid)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	if cid.Id != 0 {
		zap.L().Info("用户已经参加过该拼团！请重新选择！")
		return nil, errors.New("用户已经参加过该拼团！请重新选择！")
	}

	// 查询拼团的信息
	cidGroup, err := FindUserGroupByCid(in.Cid)
	if err != nil {
		return nil, err
	}

	// 查看拼团的人数是否已满
	if cidGroup.People >= id.People {
		zap.L().Info("拼团人数已满！请另外寻找拼团！")
		return nil, errors.New("拼团人数已满！拼团已结束！")
	}

	// 判断邀请码是否有效
	var groupBuying model_mysql.GroupBuying
	err = global.DB.Where("invitation_code =? AND cid =?", in.InvitationCode, in.Cid).First(&groupBuying).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("无效的邀请码")
		}
		return nil, err
	}

	totalPrice := float64(in.Num) * id.Price
	gb = &model_mysql.GroupBuying{
		Uid:        in.Uid,
		TotalNum:   in.Num,
		TotalPrice: totalPrice,
		Cid:        in.Cid,
		Pid:        id.ProductId,
		Price:      id.Price,
		StopTime:   id.StopTime,
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

		// 已有拼团，更新总人数，KId设为1
		result := tx.Model(&model_mysql.GroupBuying{}).
			Where("cid =?", in.Cid).
			UpdateColumn("people", gorm.Expr("people +?", 1))

		if result.Error != nil {
			return result.Error
		}

		updateResult := tx.Model(&model_mysql.GroupBuying{}).
			Where("id =?", gb.Id).
			Update("k_id", 1)

		if updateResult.Error != nil {
			return updateResult.Error
		}

		fmt.Println("已有拼团，people自增1，k_id设为1")

		// 扣减库存
		_, err := groupBuyingProduct.ReduceGroupProduct(id.ProductId, in.Num)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, errors.New("参与拼团失败")
	}

	// 重新查询最新的拼团信息
	updatedGroup, err := FindUserGroupByCid(in.Cid)
	if err != nil {
		return nil, err
	}

	if updatedGroup.People < id.People {
		zap.L().Info("当前拼团人数不足！可以分享给更多好友！")
	} else {
		// 拼团已满，可添加拼团成功逻辑
		zap.L().Info("拼团已满！恭喜成功组团！")
		UpdateGroupBuyingStatus(in.Cid, Status)

	}

	return gb, nil
}
