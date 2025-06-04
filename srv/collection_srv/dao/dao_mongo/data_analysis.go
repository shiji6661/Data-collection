package dao_mongo

import (
	"common/global"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type Message struct {
	Uid  int64 `bson:"uid"`
	Rete int64 `bson:"rete"`
}

func DataAnalysis() (int64, int64) {
	// 选择数据库和集合
	coll := global.Client.Database(global.DATABASE).Collection(global.MESSAGES_COLLECTION)
	// 构建查询条件
	filter := map[string]interface{}{
		"uid":    2,
		"status": 0,
	}
	var result Message
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			fmt.Println("未找到符合条件的记录")
		} else {
			fmt.Println("查询失败:", err)
		}
		return 0, 0
	}

	fmt.Printf("uid为2的rete值为: %d\n", result.Rete)
	return result.Rete, result.Uid

}
