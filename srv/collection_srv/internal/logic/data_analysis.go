package logic

import (
	"errors"
	"weikang/Data-collection/srv/collection_srv/dao/dao_mongo"
	"weikang/Data-collection/srv/collection_srv/proto_collection/collection"
)

// todo:数据分析
func DataAnalysis(in *collection.DataAnalysisRequest) (*collection.DataAnalysisResponse, error) {
	analysis, uid := dao_mongo.DataAnalysis()
	if analysis == 0 {
		return nil, errors.New("Data analysis failed")
	}
	return &collection.DataAnalysisResponse{
		Uid:  uid,
		Rete: analysis,
	}, nil
}
