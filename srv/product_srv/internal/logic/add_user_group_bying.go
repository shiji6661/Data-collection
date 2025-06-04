package logic

import (

	"common/global"
	user2 "common/utils/user"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"kuteng-RabbitMQ/SimlpePublish"
	"models/model_product/model_mysql"
	"order_srv/pkg"
	"product_srv/dao/dao_mysql"
	"product_srv/proto_product/product"
	"strconv"
	"way/user"
)

// todo:用户拼团
func UserCreateBuyingProduct(in *product.CreateUserGroupRequest) (*product.CreateUserGroupResponse, error) {
	// 创建拼团
	buying, err := dao_mysql.CreateGroupBuying(in)
	if err != nil {
		return nil, err
	}
	var status string
	status = strconv.FormatInt(buying.Status, 10)
	if status == "1" {
		status = "进行中"
	}
	if status == "2" {
		status = "已完成"
	}
	if status == "3" {
		status = "未完成"
	}
	products, err := dao_mysql.FindProductById(buying.Pid)
	if err != nil {
		return nil, err
	}
	ui, err := user.CommonGetUserIdInfo(in.Uid)
	if err != nil {
		return nil, err
	}
	//拼团创建成功塞入rabbitmq进行创建订单
	orderSn := uuid.New().String()
	orderMsg := map[string]interface{}{
		"OrderId":       orderSn,
		"Uid":           in.Uid,
		"RealName":      ui.UserName,
		"UserPhone":     ui.UserPhone,
		"UserAddress":   ui.UserAddress,
		"CartId":        products.CateId,
		"FreightPrice":  float64(5 * in.Num), //运费5元一件
		"TotalNum":      in.Num,
		"TotalPrice":    buying.TotalPrice,
		"PayType":       1,
		"MerId":         products.MerId,
		"ProductId":     buying.Pid,
		"CombinationId": buying.Cid,
		"PinkId":        buying.Id,
	}
	err = SimlpePublish.SimplePublishb(orderMsg)
	if err != nil {
		zap.L().Info("rabbitmq发送失败")
		return nil, errors.New("rabbitmq发送失败")
	}
	// 生成邀请码
	invitationCode, _ := user2.GenerateInviteCode(uint64(in.Uid))
	// 将邀请码保存到数据库中
	updateResult := global.DB.Model(&model_mysql.GroupBuying{}).
		Where("id =?", buying.Id).
		Update("invitation_code", invitationCode)

	if updateResult.Error != nil {
		return nil, updateResult.Error
	}
	if updateResult.RowsAffected == 0 {
		zap.L().Info("保存邀请码失败")
		return nil, errors.New("保存邀请码失败")
	}
	fmt.Println(invitationCode)
	pay := pkg.NewAliPay()
	amount := strconv.FormatFloat(buying.TotalPrice, 'f', -1, 64)
	url := pay.Pay(products.StoreName, orderSn, amount)
	return &product.CreateUserGroupResponse{
		UGroupId:     buying.Id,
		InviteCode:   invitationCode,
		NowPeopleNum: buying.People,
		Status:       status,
		Url:          url,
	}, nil
}
