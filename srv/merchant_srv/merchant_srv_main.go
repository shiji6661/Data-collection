package main

import (
	"common/initialize"
	"common/viper"
	"google.golang.org/grpc"
	"merchant_srv/grpc_merchant"
	"merchant_srv/internal"
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

	grpc_merchant.RegisterGrpc(func(server *grpc.Server) {
		internal.RegisterMerchant(server)
	})
}
