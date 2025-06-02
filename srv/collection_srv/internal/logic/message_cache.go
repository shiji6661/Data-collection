package logic

import (
	"collection_srv/dao/dao_redis"
	"collection_srv/proto_collection/collection"
	"errors"
)

// todo:redis消息缓存
func MessageCache(in *collection.MessageCacheRequest) (*collection.MessageCacheResponse, error) {
	err := dao_redis.MessageCache(in)
	if err != nil {
		return nil, errors.New("redis cache failed")
	}
	return &collection.MessageCacheResponse{Success: true}, nil

}
