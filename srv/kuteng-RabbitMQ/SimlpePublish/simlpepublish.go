package SimlpePublish

import (
	"encoding/json"
	"fmt"
	"kuteng-RabbitMQ/rabbitmq"
)

func SimplePublishb(data interface{}) error {
	mq := rabbitmq.NewRabbitMQSimple("" +
		"Order")

	marshal, err := json.Marshal(data)
	if err != nil {
		return err
	}
	mq.PublishSimple(string(marshal))
	fmt.Println("发送成功！")
	return nil
}
