package dao_mysql

import (
	"errors"
	"models/model_product/model_mysql"
	"product_srv/proto_product/product"
)

// todo 商品筛选
func ProductFilter(in *product.ProductFilterRequest) (newProduct []*product.NewProductList, err error) {
	var result []*product.NewProductList

	getProduct, err := model_mysql.GetProduct()
	if err != nil {
		return nil, errors.New("商品信息查询失败")
	}

	for _, products := range getProduct {
		// 价格区间筛选
		if in.MinPrice > 0 && products.Price < in.MinPrice {
			continue
		}
		if in.MaxPrice > 0 && products.Price > in.MaxPrice {
			continue
		}
		// 分类 ID 筛选
		if len(in.CateId) > 0 && products.CateId != in.CateId {
			continue
		}
		// 销量筛选
		if in.Sales > 0 && products.Sales < in.Sales {
			continue
		}

		// 是否热卖筛选
		if in.IsHot > 0 && products.IsHot < in.IsHot {
			continue
		}

		// 是否特价筛选
		if in.IsBenefit > 0 && products.IsBenefit < in.IsBenefit {
			continue
		}

		// 是否精品筛选
		if in.IsBest > 0 && products.IsBest < in.IsBest {
			continue
		}

		// 是否新品筛选
		if in.IsNew > 0 && products.IsNew < in.IsNew {
			continue
		}

		//// 是否包邮筛选
		//if in.Sales > 0 && products.Sales < in.Sales {
		//	continue
		//}

		// 是否秒杀状态筛选
		if in.IsSeckill > 0 && products.IsSeckill < in.IsSeckill {
			continue
		}

		// 是否砍价状态筛选
		if in.IsBargain > 0 && products.IsBargain < in.IsBargain {
			continue
		}
		result = append(result, &product.NewProductList{
			MerId:     products.MerId,
			Image:     products.Image,
			StoreName: products.StoreName,
			StoreInfo: products.StoreInfo,
			CateId:    products.CateId,
			Price:     float32(products.Price),
			VipPrice:  float32(products.VipPrice),
			Sales:     products.Sales,
			Stock:     products.Stock,
			IsPostage: products.IsPostage,
			Browse:    products.Browse,
		})
	}

	return result, nil
}
