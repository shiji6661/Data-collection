package logic

import (
	"weikang/Data-collection/srv/collection_srv/dao/dao_mongo"
	"weikang/Data-collection/srv/collection_srv/proto_collection/collection"
)

// todo:信息入库
func InformationStore(in *collection.InformationStoreRequest) (*collection.InformationStoreResponse, error) {
	err := dao_mongo.InformationStore(in)
	if err != nil {
		return nil, err
	}
	return &collection.InformationStoreResponse{Success: true}, nil
}
