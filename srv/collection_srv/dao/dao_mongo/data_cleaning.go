package dao_mongo

import (
	"common/global"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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
	// 构建聚合管道：筛选符合条件的文档ID
	// 构建聚合管道：筛选符合条件的文档ID
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"$or": []bson.M{
					{
						"CreateTime": bson.M{"$lt": nowTime},
					},
					{
						// 关键点：通过 $expr 使用聚合表达式，修正 $lt、$gt 参数传递
						"$expr": bson.M{
							"$or": []bson.M{
								{
									"$lt": []interface{}{bson.M{"$toInt": "$Heartbeat"}, 40},
								},
								{
									"$gt": []interface{}{bson.M{"$toInt": "$Heartbeat"}, 160},
								},
							},
						},
					},
				},
			},
		},
		{"$project": bson.M{"_id": 1}}, // 只返回 _id 字段
	}

	// 执行聚合查询获取待删除的文档ID
	cursor, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return fmt.Errorf("聚合查询失败: %v", err)
	}
	defer cursor.Close(context.TODO())

	var docIDs []interface{}
	for cursor.Next(context.TODO()) {
		var doc struct {
			ID interface{} `bson:"_id"`
		}
		if err := cursor.Decode(&doc); err != nil {
			return fmt.Errorf("解码文档失败: %v", err)
		}
		docIDs = append(docIDs, doc.ID)
	}

	// 批量删除符合条件的文档
	if len(docIDs) > 0 {
		result, err := coll.DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": docIDs}})
		if err != nil {
			return fmt.Errorf("删除文档失败: %v", err)
		}
		fmt.Printf("成功删除 %d 条数据\n", result.DeletedCount)
	} else {
		fmt.Println("没有找到符合条件的数据")
	}

	return nil
}
