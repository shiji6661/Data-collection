package handler

import (
	"Api/client"
	"collection_srv/proto_collection/collection"
	"context"
)

// todo: 数据采集
func DataCollection(ctx context.Context, i *collection.DataCollectionRequest) (*collection.DataCollectionResponse, error) {
	collectionClient, err := client.CollectionClient(ctx, func(ctx context.Context, in collection.CollectionClient) (interface{}, error) {
		register, err := in.DataCollection(ctx, i)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return collectionClient.(*collection.DataCollectionResponse), nil
}

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

// todo: 数据清洗
func DataCleaning(ctx context.Context, i *collection.DataCleaningRequest) (*collection.DataCleaningResponse, error) {
	collectionClient, err := client.CollectionClient(ctx, func(ctx context.Context, in collection.CollectionClient) (interface{}, error) {
		register, err := in.DataCleaning(ctx, i)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return collectionClient.(*collection.DataCleaningResponse), nil
}

// todo:数据分析
func DataAnalysis(ctx context.Context, i *collection.DataAnalysisRequest) (*collection.DataAnalysisResponse, error) {
	collectionClient, err := client.CollectionClient(ctx, func(ctx context.Context, in collection.CollectionClient) (interface{}, error) {
		register, err := in.DataAnalysis(ctx, i)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return collectionClient.(*collection.DataAnalysisResponse), nil
}

// todo:状态修改
func UpdateStatus(ctx context.Context, i *collection.UpdateStatusRequest) (*collection.UpdateStatusResponse, error) {
	collectionClient, err := client.CollectionClient(ctx, func(ctx context.Context, in collection.CollectionClient) (interface{}, error) {
		register, err := in.UpdateStatus(ctx, i)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return collectionClient.(*collection.UpdateStatusResponse), nil
}
