package logic

import (
	"collection_srv/dao/dao_mongo"
	"collection_srv/proto_collection/collection"
	"common/pkg/pkg_collection"
	"fmt"
)

// todo:数据接收
func DataCollection() (*collection.DataCollectionResponse, error) {
	// 启动MQTT客户端
	client, err := pkg_collection.StartMqtt()
	if err != nil {
		fmt.Printf("启动MQTT失败: %v\n", err)
		return nil, err
	}
	defer client.Disconnect(250)

	// 处理消息（入库逻辑）
	for msg := range pkg_collection.MessageChannel {
		// 入库处理
		err = dao_mongo.MsgCreateMongo(msg)
		if err != nil {
			fmt.Printf("入库失败: %v\n", err)
		} else {
			fmt.Printf("入库成功: [%s] %s\n", msg.Topic, string(msg.Payload))
		}
	}

	fmt.Println("程序已退出")
	return &collection.DataCollectionResponse{Success: true}, nil
}
