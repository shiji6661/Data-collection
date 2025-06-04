package internal

import (
	"google.golang.org/grpc"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
	"weikang/Data-collection/srv/product_srv/server"
)

func Register(ser *grpc.Server) {
	product.RegisterProductServer(ser, server.ServerProduct{})
}
