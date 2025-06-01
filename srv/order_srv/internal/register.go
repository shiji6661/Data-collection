package internal

import (
	"google.golang.org/grpc"
	"order_srv/proto_order/order"
	"order_srv/server"
)

func Register(ser *grpc.Server) {
	order.RegisterOrderServer(ser, server.ServerOrder{})
}
