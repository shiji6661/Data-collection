package dao_mysql

import (
	"errors"
	"fmt"
	"models/model_product/model_mysql"
	"product_srv/proto_product/product"
)

// todo:商品关键字查询
func FindProductKeyword(key string) (po []*product.NewProductList, err error) {
	keyword, err := model_mysql.FindProductKeyword(key)
	if err != nil {
		return nil, errors.New("商品数据查询失败")
	}
	if keyword == nil {
		return nil, errors.New("商品数据查询失败")
	}
	fmt.Println("11111111", keyword)
	var all []*product.NewProductList

	for _, p := range keyword {
		all = append(all, &product.NewProductList{
			MerId:     p.MerId,
			Image:     p.Image,
			StoreName: p.StoreName,
			StoreInfo: p.StoreInfo,
			CateId:    p.CateId,
			Price:     float32(p.Price),
			VipPrice:  float32(p.VipPrice),
			Sales:     p.Sales,
			Stock:     p.Stock,
			IsPostage: p.IsPostage,
			Browse:    p.Browse,
			Keyword:   p.Keyword,
		})
		fmt.Println(p.Keyword)
	}

	return all, nil
}
