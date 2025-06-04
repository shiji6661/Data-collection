package logic

import (
	"weikang/Data-collection/srv/collection_srv/dao/dao_mongo"
	"weikang/Data-collection/srv/collection_srv/proto_collection/collection"
)

// todo:数据清洗
func DataCleaning(in *collection.DataCleaningRequest) (*collection.DataCleaningResponse, error) {
	err := dao_mongo.DataCleaning()
	if err != nil {
		return nil, err
	}
	return &collection.DataCleaningResponse{Success: true}, nil
}
