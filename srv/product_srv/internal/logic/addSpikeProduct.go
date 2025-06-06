package logic

import (
	"weikang/Data-collection/srv/product_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
)

func AddSpikeProduct(in *product.AddSpikeProductRequest) (*product.AddSpikeProductResponse, error) {
	spikeProduct, err := dao_mysql.AddSpikeProduct(in)
	if err != nil {
		return nil, err
	}
	return &product.AddSpikeProductResponse{ProductId: int64(spikeProduct.ID)}, err
}
