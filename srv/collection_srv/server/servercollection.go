package server

import (
	"context"
	"weikang/Data-collection/srv/collection_srv/internal/logic"
	"weikang/Data-collection/srv/collection_srv/proto_collection/collection"
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

// todo:数据清洗
func (s ServerCollection) DataCleaning(ctx context.Context, in *collection.DataCleaningRequest) (*collection.DataCleaningResponse, error) {
	res, err := logic.DataCleaning(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// todo:数据分析
func (s ServerCollection) DataAnalysis(ctx context.Context, in *collection.DataAnalysisRequest) (*collection.DataAnalysisResponse, error) {
	res, err := logic.DataAnalysis(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// todo:状态修改
func (s ServerCollection) UpdateStatus(ctx context.Context, in *collection.UpdateStatusRequest) (*collection.UpdateStatusResponse, error) {
	res, err := logic.UpdateStatus(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
