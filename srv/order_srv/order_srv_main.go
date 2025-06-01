package main

import (
	"common/initialize"
	"common/viper"
	"google.golang.org/grpc"
	"order_srv/grpc_order"
	"order_srv/internal"
)

func main() {

	viper.InitViper()
	initialize.InitNaCos(func() {
		initialize.InitMysql()
		initialize.InitRedis()
	})
	initialize.InitConsul()
	initialize.InitMysql()
	initialize.InitRedis()
	initialize.InitEs()
	initialize.ZapInit()
	grpc_order.RegisterGrpc(func(server *grpc.Server) {
		internal.Register(server)
	})
}
