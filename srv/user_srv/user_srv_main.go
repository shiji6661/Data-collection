package main

import (
	"common/initialize"
	"common/viper"
	"google.golang.org/grpc"
	"weikang/Data-collection/srv/user_srv/grpc_user"
	"weikang/Data-collection/srv/user_srv/internal"
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
	grpc_user.RegisterGrpc(func(server *grpc.Server) {
		internal.Register(server)
	})
}
