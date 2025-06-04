package logic

import (
	"errors"
	"weikang/Data-collection/srv/collection_srv/dao/dao_redis"
	"weikang/Data-collection/srv/collection_srv/proto_collection/collection"
)

// todo:redis缓存消息查看
func GetMessageCache(in *collection.GetMessageCacheRequest) (*collection.GetMessageCacheResponse, error) {
	cache, err := dao_redis.GetMessageCache(in.Uid)
	if err != nil {
		return nil, err
	}
	//缓存消息清除
	err = dao_redis.ClearMessageCache(in.Uid)
	if err != nil {
		return nil, errors.New("缓存消息清除失败")
	}
	return &collection.GetMessageCacheResponse{
		Uid:  in.Uid,
		List: cache,
	}, nil

}
