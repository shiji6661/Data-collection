package initialize

import (
	"common/appconfig"
	"common/viper"
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type NaCosApp func()

func InitNaCos(n NaCosApp) {
	//create clientConfig
	cos := viper.Config
	clientConfig := constant.ClientConfig{
		NamespaceId:         cos.NamespaceId, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./tmp/nacos/log",
		CacheDir:            "./tmp/nacos/cache",
		LogLevel:            "debug",
	}
	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      cos.IpAddr,
			ContextPath: "/nacos",
			Port:        cos.Port,
			Scheme:      "http",
		},
	}
	// Create config client for dynamic configuration
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: cos.DataId,
		Group:  cos.Group})
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(content), &appconfig.NaCos)
	if err != nil {
		panic(err)
	}
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: cos.DataId,
		Group:  cos.Group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + cos.Group + ", dataId:" + cos.DataId + ", data:" + data)
			err = json.Unmarshal([]byte(data), &appconfig.NaCos)
			if err != nil {
				panic(err)
			}
			n()
		},
	})
	fmt.Println("nacos connect success")
}
