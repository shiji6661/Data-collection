package initialize

import (
	"common/appconfig"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
)

func InitConsul() {
	data := appconfig.NaCos
	client, err := api.NewClient(&api.Config{
		Address: fmt.Sprintf("%s:%d", data.Consul.Host, data.Consul.Port),
	})
	if err != nil {
		return
	}
	id := uuid.New().String()
	err = client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      id,
		Name:    data.Consul.Name,
		Tags:    []string{"GRPC"},
		Port:    data.Group.Port,
		Address: data.Group.Host,
		Check: &api.AgentServiceCheck{
			Interval:                       "5s",
			Timeout:                        "5s",
			GRPC:                           fmt.Sprintf("%s:%d", data.Group.Host, data.Group.Port),
			DeregisterCriticalServiceAfter: "30s",
		},
	})
	if err != nil {
		return
	}
	fmt.Println("consul connect success")
}
