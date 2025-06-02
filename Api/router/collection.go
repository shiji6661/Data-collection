package router

import (
	"Api/internal/trigger"
	"Api/pkg"
	"github.com/gin-gonic/gin"
)

func LoadCollection(r *gin.Engine) {
	collection := r.Group("/collection")
	{
		collection.Use(pkg.JWTAuth("2209A"))
		collection.GET("/ws", trigger.Echo) //TODO:ws入口
	}
}
