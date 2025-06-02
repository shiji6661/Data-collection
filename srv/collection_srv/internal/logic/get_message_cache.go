package logic

import (
	"collection_srv/dao/dao_redis"
	"collection_srv/proto_collection/collection"
	"errors"
)

// todo:redis缓存消息查看
func GetMessageCache(in *collection.GetMessageCacheRequest) (*collection.GetMessageCacheResponse, error) {
	cache, err := dao_redis.GetMessageCache(in.Uid)
	if err != nil {
		return nil, err
	}
	if cache == 0 {
		return nil, errors.New("no cache")
	}
	return &collection.GetMessageCacheResponse{
		Uid:       in.Uid,
		Heartbeat: cache,
	}, nil

}
