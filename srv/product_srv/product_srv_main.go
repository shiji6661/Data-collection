package main

import (
	"common/initialize"
	"common/viper"
	"google.golang.org/grpc"
	"product_srv/grpc_product"
	"product_srv/internal"
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
	initialize.InitMongoDB()
	initialize.ZapInit()
	//global.DB.AutoMigrate(&model_mysql.GroupBuying{})
	//dao_redis.SyncProductsToRedis(context.Background(), 4, 200)
	grpc_product.RegisterGrpc(func(server *grpc.Server) {
		internal.Register(server)
	})
}
