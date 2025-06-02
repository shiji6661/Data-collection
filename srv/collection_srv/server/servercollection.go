package server

import (
	"collection_srv/internal/logic"
	"collection_srv/proto_collection/collection"
	"context"
)

type ServerCollection struct {
	collection.UnimplementedCollectionServer
}

// todo:数据接收
func (s ServerCollection) DataCollection(ctx context.Context, in *collection.DataCollectionRequest) (*collection.DataCollectionResponse, error) {
	dataCollection, err := logic.DataCollection()
	if err != nil {
		return nil, err
	}
	return dataCollection, nil
}

// todo:信息入库
func (s ServerCollection) InformationStore(ctx context.Context, in *collection.InformationStoreRequest) (*collection.InformationStoreResponse, error) {
	res, err := logic.InformationStore(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// todo:redis消息缓存
func (s ServerCollection) MessageCache(ctx context.Context, in *collection.MessageCacheRequest) (*collection.MessageCacheResponse, error) {
	res, err := logic.MessageCache(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// todo:redis缓存消息查看
func (s ServerCollection) GetMessageCache(ctx context.Context, in *collection.GetMessageCacheRequest) (*collection.GetMessageCacheResponse, error) {
	res, err := logic.GetMessageCache(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
