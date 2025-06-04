package logic

import (
	"product_srv/pkg"
	"product_srv/proto_product/product"
)

// 将商品信息写入ES
func ProductCreateToES(in *product.ProductCreateToESRequest) (*product.ProductCreateToESResponse, error) {
	pkg.ProductCreateToES(in.Table)
	return &product.ProductCreateToESResponse{Success: true}, nil
}
