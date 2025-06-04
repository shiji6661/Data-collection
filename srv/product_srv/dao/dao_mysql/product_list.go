package dao_mysql

import (
	"errors"
	"go.uber.org/zap"
	"models/model_product/model_mysql"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
)

// 商品列表展示
func ProductList() ([]*product.ProductItem, error) {
	pro := &model_mysql.Product{}
	list, err := pro.ProductList()
	if err != nil {
		zap.L().Info("查询失败")
		return nil, errors.New("查询失败！")
	}
	var productList []*product.ProductItem
	for _, p := range list {
		productList = append(productList, &product.ProductItem{
			StoreName: p.StoreName,
			StoreInfo: p.StoreInfo,
			CateId:    p.CateId,
			Price:     float32(p.Price),
			Sales:     p.Sales,
			Browse:    p.Browse,
		})
	}
	return productList, nil
}
