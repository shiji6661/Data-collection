package logic

import (
	"weikang/Data-collection/srv/collection_srv/dao/dao_redis"
	"weikang/Data-collection/srv/collection_srv/proto_collection/collection"
)

// todo:redis消息缓存
func MessageCache(in *collection.MessageCacheRequest) (*collection.MessageCacheResponse, error) {
	err := dao_redis.MessageCache(in.Uid, in.Tid, in.Heartbeat)
	if err != nil {
		return nil, err
	}
	return &collection.MessageCacheResponse{Success: true}, nil

}
