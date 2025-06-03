package logic

import (
	"collection_srv/dao/dao_mongo"
	"collection_srv/proto_collection/collection"
	"fmt"
)

// mongo
// todo:状态修改
func UpdateStatus(in *collection.UpdateStatusRequest) (*collection.UpdateStatusResponse, error) {

	fmt.Println("uid", in.Uid)
	fmt.Println("rete", in.Rete)
	err := dao_mongo.UpdateStatus(in.Uid, in.Rete)
	if err != nil {
		return nil, err
	}
	return &collection.UpdateStatusResponse{Success: true}, nil

}
