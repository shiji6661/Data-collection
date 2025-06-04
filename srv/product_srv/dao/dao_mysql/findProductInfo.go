package dao_mysql

import (
	"errors"
	"fmt"
	"models/model_product/model_mysql"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
)

// todo 商品详情
func FindProductInfo(name string) (Product []*product.ProductList, err error) {

	info, err := model_mysql.FindProductInfo(name)
	if err != nil {
		return nil, errors.New("商品查询失败")
	}
	if info == nil {
		return nil, errors.New("商品信息查询失败")
	}

	var all []*product.ProductList
	for _, list := range info {
		all = append(all, &product.ProductList{
			Id:            list.Id,
			MerId:         list.MerId,
			StoreName:     list.StoreName,
			StoreInfo:     list.StoreInfo,
			CateId:        list.CateId,
			Price:         float32(list.Price),
			Sales:         list.Sales,
			Stock:         list.Stock,
			IsShow:        list.IsShow,
			Cost:          float32(list.Cost),
			IsGood:        list.IsGood,
			Browse:        list.Browse,
			ProductSortId: list.ProductSortId,
			Pid:           list.Pid,
			CateName:      list.CateName,
			Sort:          list.Sort,
		})
		fmt.Println(list.StoreName)
		fmt.Println(list.CateName)

	}

	return all, nil
}
