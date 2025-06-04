package trigger

import (
	"Api/global"
	"Api/internal/handler"
	"Api/internal/response"
	"collection_srv/proto_collection/collection"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

// TODO:数据采集
func DataCollection(c *gin.Context) {
	coll, err := handler.DataCollection(c, &collection.DataCollectionRequest{})
	if err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	response.ResponseSuccess(c, coll)
}

// TODO: 数据清洗
func DataCleaning(c *gin.Context) {
	cleaning, err := handler.DataCleaning(c, &collection.DataCleaningRequest{})
	if err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	response.ResponseSuccess(c, cleaning)
}

// 模拟,应为自动监控警报
// TODO: 数据分析
func DataAnalysis(c *gin.Context) {
	//数据查询
	analysis, err := handler.DataAnalysis(c, &collection.DataAnalysisRequest{})
	if err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	fmt.Println("心跳值:", analysis.Rete)
	if analysis.Rete >= global.HEARTBEAT_BIG {
		// 构建消息内容
		msg := fmt.Sprintf("警告：您的心跳值 %d,已超过阈值 %d",
			analysis.Rete, global.HEARTBEAT_BIG)

		// 获取WebSocket连接
		wsConn, exists := OnlineUser[analysis.Uid]
		if !exists {
			response.ResponseError400(c, "用户未建立WebSocket连接")
			return
		}

		// 发送消息（直接调用响应函数或复用SendFunc）
		err = response.WsResponseSuccess(wsConn, msg)
		if err != nil {
			log.Printf("向用户 %d 发送心跳警告消息失败: %v", analysis.Uid, err)
		} else {
			//发送成功后，更新状态
			status, err := handler.UpdateStatus(c, &collection.UpdateStatusRequest{
				Uid:  analysis.Uid,
				Rete: analysis.Rete,
			})
			if err != nil {
				response.ResponseError400(c, err.Error())
				return
			}
			fmt.Println("状态修改成功", status)
			//response.ResponseSuccess(c, status.Success)
		}

	}
	response.ResponseSuccess(c, analysis.Rete)
}
