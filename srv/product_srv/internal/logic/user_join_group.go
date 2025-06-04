package logic

import (
	"Data-collection/way/user"
	"errors"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"kuteng-RabbitMQ/SimlpePublish"
	"order_srv/pkg"
	"weikang/Data-collection/srv/product_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/product_srv/proto_product/product"

	"strconv"
)

// todo:用户加入拼团
func UserJoinGroup(in *product.UserJoinGroupRequest) (*product.UserJoinGroupResponse, error) {
	gb, err := dao_mysql.JoinGroupBuying(in, in.InvitationCode)
	if err != nil {
		return nil, err
	}
	var status string
	status = strconv.FormatInt(gb.Status, 10)
	if status == "1" {
		status = "进行中"
	}
	if status == "2" {
		status = "已完成"
	}
	if status == "3" {
		status = "未完成"
	}
	products, err := dao_mysql.FindProductById(gb.Pid)
	if err != nil {
		return nil, err
	}
	ui, err := user.CommonGetUserIdInfo(in.Uid)
	if err != nil {
		return nil, err
	}
	//支付链接
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
		"TotalPrice":    gb.TotalPrice,
		"PayType":       1,
		"MerId":         products.MerId,
		"ProductId":     gb.Pid,
		"CombinationId": gb.Cid,
		"PinkId":        gb.Id,
	}

	err = SimlpePublish.SimplePublishb(orderMsg)
	if err != nil {
		zap.L().Info("rabbitmq发送失败")
		return nil, errors.New("rabbitmq发送失败")
	}
	pay := pkg.NewAliPay()
	amount := strconv.FormatFloat(gb.TotalPrice, 'f', -1, 64)
	url := pay.Pay(products.StoreName, orderSn, amount)
	return &product.UserJoinGroupResponse{
		UGroupId:    gb.Id,
		ProductName: products.StoreName,
		Status:      status,
		Url:         url,
	}, nil
}
