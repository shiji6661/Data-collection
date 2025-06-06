package logic

import (
	"weikang/Data-collection/srv/product_srv/pkg"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
)

// todo 查询ES中的商品信息
func ProductSearchToEs(in *product.ProductSearchESRequest) (*product.ProductSearchESResponse, error) {
	es, err := pkg.ProductSearchToEs(in.Name)
	if err != nil {
		return nil, err
	}
	return &product.ProductSearchESResponse{
		List: es,
	}, nil
}
