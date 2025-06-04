package viper

import (
	"fmt"
	"github.com/spf13/viper"
)

type AppConfig struct {
	NaCosAppFig
	MqttFig
}
type NaCosAppFig struct {
	NamespaceId string
	IpAddr      string
	Port        uint64
	DataId      string
	Group       string
}

type MqttFig struct {
	Broker   string
	Topic    string
	Qos      int64
	UserName string
	Password string
}

var Config AppConfig

func InitViper() {
	viper.SetConfigFile("../../common/appconfig/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
	fmt.Println("viper连接成功")
}
