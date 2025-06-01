package logic

import (
	"errors"
	mysql2 "models/model_product/model_mysql"
	"order_srv/dao/dao_mysql"
	"order_srv/pkg"
	"order_srv/proto_order/order"
	"strconv"
)

func CreateOrderBargain(in *order.CreateOrderBargainRequest) (*order.CreateOrderBargainResponse, error) {
	b := mysql2.Bargain{}
	id, err := b.FindBargainProductById(int(in.BargainId))
	if err != nil {
		return nil, err
	}
	if id.ID == 0 {
		return nil, errors.New("砍价商品不存在")
	}
	c := mysql2.BargainUser{}
	userId, err := c.FindBargainUserId(int(in.UserId))
	if err != nil {
		return nil, err
	}
	if userId.Id == 0 {
		return nil, errors.New("未发起砍价")
	}
	bargain, err := dao_mysql.CreateOrderBargain(in)
	if err != nil {
		return nil, err
	}
	pay := pkg.NewAliPay()
	Price := strconv.FormatFloat(userId.Price, 'f', 2, 64)
	s := pay.Pay(id.StoreName, bargain.OrderId, Price)
	return &order.CreateOrderBargainResponse{
		OrderId: bargain.Id,
		Url:     s,
	}, err
}
