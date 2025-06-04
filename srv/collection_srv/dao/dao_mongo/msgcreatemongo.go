package dao_mongo

import (
	"common/global"
	"common/pkg/pkg_collection"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

type PayloadStruct struct {
	Uid  int `json:"uid"`
	Rete int `json:"rete"`
}

func MsgCreateMongo(msg pkg_collection.MqttMessage) error {
	if global.Client == nil {
		return errors.New("MongoDB client is not initialized")
	}

	// 解析JSON
	var payload PayloadStruct
	err := json.Unmarshal(msg.Payload, &payload)
	if err != nil {
		return fmt.Errorf("JSON解析失败: %v", err)
	}
	// 选择数据库和集合
	collection := global.Client.Database(global.DATABASE).Collection(global.MESSAGES_COLLECTION)
	now := msg.Time.Format("20060102")

	// 插入消息
	d := bson.D{
		{Key: "topic", Value: msg.Topic},
		{Key: "uid", Value: payload.Uid},
		{Key: "rete", Value: payload.Rete},
		{Key: "time", Value: now},
		{Key: "status", Value: 0},
	}
	_, err = collection.InsertOne(global.CTX, d)
	if err != nil {
		return err
	}
	return nil
}
