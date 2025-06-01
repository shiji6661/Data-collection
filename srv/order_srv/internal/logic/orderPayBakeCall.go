package logic

import (
	"errors"
	"order_srv/dao/dao_mysql"
	"order_srv/proto_order/order"
)

func OrderPayBakeCall(in *order.OrderPayBakeCallRequest) (*order.OrderPayBakeCallResponse, error) {
	err := dao_mysql.OrderBack(in.OrderId, in.Status)
	if err != nil {
		return nil, errors.New("支付状态修改失败")
	}
	return &order.OrderPayBakeCallResponse{Success: true}, nil
}
