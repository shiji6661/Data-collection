package dao_mysql

import (
	"errors"
	"models/model_product/model_mysql"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
)

// todo:商品推荐
func FindProductRecommend() (po []*product.NewProductList, err error) {
	products, err := model_mysql.ProductRecommend()
	if err != nil {
		return nil, errors.New("商品查询失败")
	}
	if products == nil {
		return nil, errors.New("商品信息查询失败")
	}

	var all []*product.NewProductList

	for _, p := range products {
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
		})
	}

	return all, nil
}
