package dao_mysql

import (
	"kuteng-RabbitMQ/SimpleRecieve"
	"order_srv/proto_order/order"
)

func OrderCreate(in *order.OrderCreateRequest) (*order.OrderCreateResponse, error) {
	SimpleRecieve.SimpleReceive()
	return &order.OrderCreateResponse{Success: "ok"}, nil
}
