package logic

import (
	"weikang/Data-collection/srv/product_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
)

// todo 商品筛选
func ProductFilter(in *product.ProductFilterRequest) (*product.ProductFilterResponse, error) {
	list, err := dao_mysql.ProductFilter(in)
	if err != nil {
		return nil, err
	}
	return &product.ProductFilterResponse{List: list}, nil
}
