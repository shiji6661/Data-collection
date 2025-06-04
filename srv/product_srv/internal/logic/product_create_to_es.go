package logic

import (
	"weikang/Data-collection/srv/product_srv/pkg"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
)

// 将商品信息写入ES
func ProductCreateToES(in *product.ProductCreateToESRequest) (*product.ProductCreateToESResponse, error) {
	pkg.ProductCreateToES(in.Table)
	return &product.ProductCreateToESResponse{Success: true}, nil
}
