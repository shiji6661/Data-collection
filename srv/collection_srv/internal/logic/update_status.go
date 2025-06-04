package logic

import (
	"collection_srv/dao/dao_mongo"
	"collection_srv/proto_collection/collection"
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
