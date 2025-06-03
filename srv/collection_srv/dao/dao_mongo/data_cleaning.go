package dao_mongo

import (
	"common/global"
	"context"
	"fmt"
	"time"
)

func DataCleaning() error {
	// 获取当前时间
	now := time.Now()
	// 计算七天前的时间
	sevenDaysAgo := now.Add(-7 * 24 * time.Hour)
	nowTime := sevenDaysAgo.Format("20060102")

	// 选择数据库和集合
	coll := global.Client.Database(global.DATABASE).Collection(global.PRIVATE)
	// 构建删除条件，这里假设 "CreateTime" 字段是时间类型
	filter := map[string]interface{}{
		"CreateTime": map[string]interface{}{
			"$lt": nowTime,
		},
	}

	// 执行删除操作
	result, err := coll.DeleteMany(context.TODO(), filter)
	if err != nil {
		fmt.Println("删除数据失败:", err)
		return err
	}
	fmt.Printf("成功删除 %d 条数据\n", result.DeletedCount)
	return nil
}
