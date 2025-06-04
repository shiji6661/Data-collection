package logic

import (
	"weikang/Data-collection/srv/product_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
)

// todo: 商品展示
func ProductList(in *product.ProductListRequest) (*product.ProductListResponse, error) {
	list, err := dao_mysql.ProductList()
	if err != nil {
		return nil, err
	}
	return &product.ProductListResponse{List: list}, nil
}
