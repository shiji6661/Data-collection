package handler

import (
	"Api/client"
	"collection_srv/proto_collection/collection"
	"context"
)

// todo: 信息入库
func InformationStore(ctx context.Context, i *collection.InformationStoreRequest) (*collection.InformationStoreResponse, error) {
	collectionClient, err := client.CollectionClient(ctx, func(ctx context.Context, in collection.CollectionClient) (interface{}, error) {
		register, err := in.InformationStore(ctx, i)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return collectionClient.(*collection.InformationStoreResponse), nil
}

// todo: redis消息缓存
func MessageCache(ctx context.Context, i *collection.MessageCacheRequest) (*collection.MessageCacheResponse, error) {
	collectionClient, err := client.CollectionClient(ctx, func(ctx context.Context, in collection.CollectionClient) (interface{}, error) {
		register, err := in.MessageCache(ctx, i)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return collectionClient.(*collection.MessageCacheResponse), nil
}

// todo: redis缓存消息查看
func GetMessageCache(ctx context.Context, i *collection.GetMessageCacheRequest) (*collection.GetMessageCacheResponse, error) {
	collectionClient, err := client.CollectionClient(ctx, func(ctx context.Context, in collection.CollectionClient) (interface{}, error) {
		register, err := in.GetMessageCache(ctx, i)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return collectionClient.(*collection.GetMessageCacheResponse), nil
}
