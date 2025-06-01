package router

import (
	"Api/internal/middleware"
	"Api/internal/trigger"
	"Api/pkg"
	"github.com/gin-gonic/gin"
)

func LoadOrder(r *gin.Engine) {
	r.Use(middleware.Logger())
	order := r.Group("/order")
	{
		order.POST("/call/back", trigger.OrderCallBack) //todo 支付回调

		order.POST("/create/order/message", trigger.RabbitMqToCreate) //todo 消费队列进行下单

		order.Use(pkg.JWTAuth("2209A"))
		order.POST("/create/order", trigger.CreateOrder)          //todo 商品下单
		order.POST("/create/order/cart", trigger.CreateOrderCart) //todo 商品下单(购物车)
		order.POST("/order/create/spike", trigger.OrderCreateSpike)
		order.POST("/order/pay", trigger.OrderPay) //todo 订单支付(异步)
		order.POST("/create/order/bargain", trigger.CreateOrderBargain)

		order.POST("/pay", trigger.OrderPay)         //todo 订单支付(异步)
		order.POST("/qr/code", trigger.CreateOrCode) //todo 生成二维码

	}
	Merchant := r.Group("/product")
	{
		Merchant.Use(pkg.JWTAuth("Merchant"))
		Merchant.POST("/count", trigger.MerchantCountAmount) //todo 商家查看销售额
	}
}
