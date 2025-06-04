package dao_mysql

import (
	"errors"
	"models/model_product/model_mysql"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
)

// todo:商品分类浏览
func ProductCategory(str string) (newProduct []*product.ProductList, err error) {
	products, err := model_mysql.ProductCategory(str)
	if err != nil {
		return nil, errors.New("商品数据查询失败")
	}
	if products == nil {
		return nil, errors.New("商品信息查询失败")
	}

	var all []*product.ProductList
	for _, list := range products {
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

	}

	return all, nil
}
