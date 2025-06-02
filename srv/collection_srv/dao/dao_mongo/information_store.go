package dao_mongo

import (
	"collection_srv/proto_collection/collection"
	"common/global"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// todo:信息入库
func InformationStore(in *collection.InformationStoreRequest) error {
	// 选择数据库和集合
	coll := global.Client.Database(in.Database).Collection(in.Table)
	bsd := bson.D{
		{"Uid", in.Uid},
		{"Tid", in.Tid},
		{"Heartbeat", in.Heartbeat},
		{"CreateTime", time.Now().String()},
	}
	_, err := coll.InsertOne(global.CTX, bsd)
	if err != nil {
		return err
	}
	return nil
}
