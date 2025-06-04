package logic

import (
	"collection_srv/dao/dao_mongo"
	"collection_srv/proto_collection/collection"
	"errors"

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
