package router

import (
	"Api/internal/trigger"
	"Api/pkg"
	"github.com/gin-gonic/gin"
)

func LoadCollection(r *gin.Engine) {
	collection := r.Group("/collection")
	{
		collection.POST("/data/collection", trigger.DataCollection) // TODO: 数据采集
		collection.Use(pkg.JWTAuth("2209A"))
		collection.GET("/ws", trigger.Echo)                     //TODO:ws入口
		collection.POST("/data/cleaning", trigger.DataCleaning) // TODO: 数据清洗
		collection.POST("/data/analysis", trigger.DataAnalysis) // TODO: 数据分析

	}
}
