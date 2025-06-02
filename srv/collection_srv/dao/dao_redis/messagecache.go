package dao_redis

import (
	"collection_srv/proto_collection/collection"
	"common/global"
	"fmt"
)

func MessageCache(in *collection.MessageCacheRequest) error {
	key := fmt.Sprintf("%d:%d", in.Tid, in.Uid)
	return global.Rdb.Set(global.CTX, key, in.Heartbeat, 0).Err()
}

func GetMessageCache(id int64) (int64, error) {
	key := fmt.Sprintf("%d:*", id)
	return global.Rdb.Get(global.CTX, key).Int64()
}
