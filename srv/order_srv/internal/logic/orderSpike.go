package logic

import (
	"order_srv/dao/dao_mysql"
	"order_srv/proto_order/order"
)

func OrderCreate(in *order.OrderCreateRequest) (*order.OrderCreateResponse, error) {
	create, err := dao_mysql.OrderCreate(in)
	if err != nil {
		return nil, err
	}
	return &order.OrderCreateResponse{Success: create.Success}, err
}
