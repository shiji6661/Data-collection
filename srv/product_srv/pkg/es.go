package pkg

import (
	"bytes"
	"common/global"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"models/model_product/model_mysql"
	"product_srv/proto_product/product"
	"strconv"
	"sync"
)

func ProductCreateToES(table string) {
	var (
		p  []*model_mysql.Product
		wg sync.WaitGroup
	)

	global.DB.Find(&p)

	for _, title := range p {
		wg.Add(1)

		go func(p *model_mysql.Product) {
			defer wg.Done()

			// Build the request body.
			data, err := json.Marshal(p)
			if err != nil {
				log.Fatalf("Error marshaling document: %s", err)
			}

			// Set up the request object.
			req := esapi.IndexRequest{
				Index:      table,
				DocumentID: strconv.Itoa(int(p.ID)),
				Body:       bytes.NewReader(data),
				Refresh:    "true",
			}

			// Perform the request with the client.
			res, err := req.Do(context.Background(), global.Es)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()
		}(title)
	}
	wg.Wait()

}

func ProductSearchToEs(name string) ([]*product.NewProductList, error) {
	var buf bytes.Buffer
	query := map[string]interface{}{}
	if name == "" {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"match_all": map[string]interface{}{},
			},
		}
	} else {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"match": map[string]interface{}{
					"store_name": name,
				},
			},
		}
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := global.Es.Search(
		global.Es.Search.WithContext(context.Background()),
		global.Es.Search.WithIndex("product"),
		global.Es.Search.WithBody(&buf),
		global.Es.Search.WithTrackTotalHits(true),
		global.Es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	var r map[string]interface{}

	json.NewDecoder(res.Body).Decode(&r)

	var All []*product.NewProductList

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		Source := hit.(map[string]interface{})["_source"].(map[string]interface{})

		goods := &product.NewProductList{
			MerId:     int64(Source["mer_id"].(float64)),
			Image:     Source["image"].(string),
			StoreName: Source["store_name"].(string),
			StoreInfo: Source["store_info"].(string),
			CateId:    Source["cate_id"].(string),
			Price:     float32(Source["price"].(float64)),
			VipPrice:  float32(Source["vip_price"].(float64)),
			Sales:     int64(Source["sales"].(float64)),
			Stock:     int64(Source["stock"].(float64)),
			IsPostage: int64(Source["is_postage"].(float64)),
			Browse:    int64(Source["browse"].(float64)),
		}
		All = append(All, goods)
	}
	return All, nil
}
