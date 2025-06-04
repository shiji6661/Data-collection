package main

import (
	"collection_srv/dao/dao_mongo"
	"common/initialize"
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	initialize.InitMongoDB()
	// 创建一个定时任务调度器
	c := cron.New(cron.WithSeconds())
	// 添加定时任务，每天凌晨1点执行数据清洗
	_, err := c.AddFunc("0 0 1 * * ?", func() {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		// 执行数据清洗逻辑
		fmt.Println("数据清洗开始")
		// TODO: 数据清洗
		err := dao_mongo.DataCleaning()
		if err != nil {
			fmt.Printf("数据清洗失败: %v\n", err)
			return
		}

	})
	if err != nil {
		return
	}
	c.Start()
	select {}
}
