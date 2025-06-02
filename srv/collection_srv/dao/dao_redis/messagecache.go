package dao_redis

import (
	"collection_srv/proto_collection/collection"
	"common/global"
	"fmt"
)

func MessageCache(uid, tid int64, heartbeat string) error {
	key := fmt.Sprintf("%d:%d", tid, uid)
	return global.Rdb.LPush(global.CTX, key, heartbeat).Err()
}

func GetMessageCache(id int64) (list []*collection.GetMessageCache, err error) {
	key := fmt.Sprintf("%d:%d", id, global.MANAGE_ID)
	lenCmd := global.Rdb.LLen(global.CTX, key)
	fmt.Printf("获取到的列表长度命令结果: %v，错误: %v\n", lenCmd.Val(), lenCmd.Err())
	if lenCmd.Err() != nil {
		fmt.Println("获取列表长度失败:", lenCmd.Err())
		return
	}
	val := lenCmd.Val()
	result := global.Rdb.LRange(global.CTX, key, 0, val-1).Val()
	fmt.Printf("LRange 获取结果: %v，错误: %v\n", result, lenCmd.Err())
	fmt.Println(result)
	for _, s := range result {
		list = append(list, &collection.GetMessageCache{
			Heartbeat: s,
		})
	}
	return list, nil
}

func ClearMessageCache(uid int64) error {
	return global.Rdb.Del(global.CTX, fmt.Sprintf("%d:%d", uid, global.MANAGE_ID)).Err()
}
