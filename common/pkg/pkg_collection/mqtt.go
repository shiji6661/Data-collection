package pkg_collection

import (
	"common/global"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 配置信息

// 消息结构
type MqttMessage struct {
	Topic   string
	Payload []byte
	Time    time.Time
	QoS     byte
}

// 消息通道
var MessageChannel = make(chan MqttMessage)

// 启动MQTT客户端
func StartMqtt() (mqtt.Client, error) {
	// 创建MQTT客户端选项
	clientID := fmt.Sprintf("go-mqtt-client-%d", rand.Intn(10000))
	opts := mqtt.NewClientOptions()
	opts.AddBroker(global.BROKER)
	opts.SetClientID(clientID)
	opts.SetUsername(global.USER_NAME)
	opts.SetPassword(global.PASSWORD)
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(10 * time.Second)
	opts.SetOnConnectHandler(func(client mqtt.Client) {
		fmt.Println("成功连接到MQTT服务器")
		if token := client.Subscribe(global.TOPIC, global.QOS, nil); token.Wait() && token.Error() != nil {
			log.Fatalf("订阅主题失败: %v", token.Error())
		}
		fmt.Printf("已订阅主题: %s\n", global.TOPIC)
	})

	// 消息处理 - 将消息发送到通道
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		MessageChannel <- MqttMessage{
			Topic:   msg.Topic(),
			Payload: msg.Payload(),
			Time:    time.Now(),
			QoS:     msg.Qos(),
		}
	})

	client := mqtt.NewClient(opts)
	token := client.Connect()
	token.Wait()

	if token.Error() != nil {
		// 打印详细的错误信息
		fmt.Printf("MQTT连接失败: 错误码=%d, 错误信息=%v\n", token.(*mqtt.ConnectToken).ReturnCode(), token.Error())
		return nil, fmt.Errorf("连接MQTT服务器失败: %w", token.Error())
	}
	// 启动信号监听协程
	go handleInterrupt(client)

	return client, nil
}

// 处理中断信号
func handleInterrupt(client mqtt.Client) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt

	fmt.Println("\n接收到中断信号，正在关闭...")
	client.Disconnect(250)
	close(MessageChannel)
}

// 发布消息
//func PublishMessage(message string) error {
//	// 这里需要持有client引用，实际使用中应通过参数或全局变量传递
//	// 示例代码，不可直接运行
//	return fmt.Errorf("PublishMessage未实现，需要传递MQTT客户端引用")
//}
