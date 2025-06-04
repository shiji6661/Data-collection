package dao_mongo

import (
	"common/global"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"weikang/Data-collection/srv/collection_srv/proto_collection/collection"
)

// todo:信息入库
func InformationStore(in *collection.InformationStoreRequest) error {
	// 选择数据库和集合
	coll := global.Client.Database(in.Database).Collection(in.Table)
	now := time.Now().Format("20060102")
	bsd := bson.D{
		{Key: "Uid", Value: in.Uid},
		{Key: "Tid", Value: in.Tid},
		{Key: "Heartbeat", Value: in.Heartbeat},
		{Key: "CreateTime", Value: now},
	}
	_, err := coll.InsertOne(global.CTX, bsd)
	if err != nil {
		return err
	}
	return nil
}
