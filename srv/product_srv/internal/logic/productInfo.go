package logic

import (
	"weikang/Data-collection/srv/product_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
)

// todo:商品详情逻辑
func ProductInfo(in *product.ProductInfoRequest) (*product.ProductInfoResponse, error) {
	all, err := dao_mysql.FindProductInfo(in.StoreName)
	if err != nil {
		return nil, err
	}

	return &product.ProductInfoResponse{
		List: all,
	}, nil
}
