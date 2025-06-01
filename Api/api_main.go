package main

import (
	"Api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.LoadUser(r)     //todo:用户组路由
	router.LoadMerchant(r) //todo:商家路由组
	router.LoadProduct(r)  //todo:商品路由组
	router.LoadOrder(r)    //todo:订单路由组
	router.LoadArticle(r)  //todo:文章路由组
	r.Run(":8888")         // 监听并在 0.0.0.0:8080 上启动服务
}
