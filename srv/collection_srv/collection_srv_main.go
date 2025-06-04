package main

import (
	"common/initialize"
	"common/viper"
	"google.golang.org/grpc"
	"weikang/Data-collection/srv/collection_srv/grpc_collection"
	"weikang/Data-collection/srv/collection_srv/internal"
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
