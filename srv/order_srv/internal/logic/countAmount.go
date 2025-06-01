package logic

import (
	"models/model_order/model_mysql"
	"order_srv/proto_order/order"
)

func MerchantCountAmount(in *order.MerchantCountAmountRequest) (*order.MerchantCountAmountResponse, error) {
	o := model_mysql.Order{}
	amount, err := o.CountAmount(int(in.MerId))
	if err != nil {
		return nil, err
	}
	return &order.MerchantCountAmountResponse{
		Amount: float32(amount),
	}, nil
}
