package internal

import (
	"google.golang.org/grpc"
	"product_srv/proto_product/product"
	"product_srv/server"
)

func Register(ser *grpc.Server) {
	product.RegisterProductServer(ser, server.ServerProduct{})
}
