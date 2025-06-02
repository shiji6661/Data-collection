package main

import (
	"collection_srv/grpc_collection"
	"collection_srv/internal"
	"common/initialize"
	"common/viper"
	"google.golang.org/grpc"
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
	initialize.InitMongoDB()

	grpc_collection.RegisterGrpc(func(server *grpc.Server) {
		internal.Register(server)
	})
}
