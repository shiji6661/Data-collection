package main

import (
	"article_srv/grpc_article"
	"article_srv/internal"
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

	grpc_article.RegisterGrpc(func(server *grpc.Server) {
		internal.RegisterArticle(server)
	})
}
