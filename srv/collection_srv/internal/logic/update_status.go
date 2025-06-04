package logic

import (
	"weikang/Data-collection/srv/collection_srv/dao/dao_mongo"
	"weikang/Data-collection/srv/collection_srv/proto_collection/collection"
)

// mongo
// todo:状态修改
func UpdateStatus(in *collection.UpdateStatusRequest) (*collection.UpdateStatusResponse, error) {

	err := dao_mongo.UpdateStatus(in.Uid, in.Rete)
	if err != nil {
		return nil, err
	}
	return &collection.UpdateStatusResponse{Success: true}, nil

}
