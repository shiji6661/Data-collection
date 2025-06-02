package initialize

import (
	"common/global"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
)

func InitEs() {
	var err error
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://14.103.140.251:9200",
		},
		// ...
	}
	global.Es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println("elasticsearch connect success")
}
