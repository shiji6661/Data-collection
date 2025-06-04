package dao_mongo

import (
	"common/global"
	"context"
	"errors"
	"fmt"
)

func UpdateStatus(uid, rete int64) error {
	// 选择数据库和集合
	coll := global.Client.Database(global.DATABASE).Collection(global.MESSAGES_COLLECTION)
	// 查询条件
	filter := map[string]interface{}{
		"uid":  uid,
		"rete": rete,
	}
	// 更新操作
	update := map[string]interface{}{
		"$set": map[string]interface{}{
			"status": 1,
		},
	}
	// 执行更新
	result, err := coll.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return errors.New("更新数据失败")
	}
	fmt.Printf("匹配到 %d 条数据，成功更新 %d 条数据\n", result.MatchedCount, result.ModifiedCount)
	return nil
}
