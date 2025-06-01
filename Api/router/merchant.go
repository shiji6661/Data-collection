package router

import (
	"Api/internal/middleware"
	"Api/internal/trigger"
	"Api/pkg"
	"github.com/gin-gonic/gin"
)

func LoadMerchant(c *gin.Engine) {
	c.Use(middleware.Logger())
	merchant := c.Group("./merchant")
	{
		merchant.POST("/register", trigger.MerchantRegister) // todo:注册

		merchant.POST("/login", trigger.MerchantLogin) // todo:登录

		merchant.POST("/send/sms", trigger.MerchantSms) // todo:商家端短信发送

		merchant.POST("/send/email", trigger.MerchantEmail) // todo:商家端邮件发送

		merchant.Use(pkg.JWTAuth("Merchant"))

		merchant.POST("/parse/code", trigger.MerchantParse) // todo:核销二维码

		merchant.POST("/statistics/store", trigger.StatisticsStoreDailyData) // TODO 按日统计店铺数据

	}
}
