package logic

import (
	"kuteng-RabbitMQ/SimpleRecieve"
	"order_srv/proto_order/order"
)

// 消费队列进行下单
func RabbitMqToCreate(in *order.CreateOrderMessageRequest) (*order.CreateOrderMessageResponse, error) {
	err := SimpleRecieve.SimpleReceive()
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderMessageResponse{
		Ping: "创建订单成功",
	}, nil
}
