package logic

import (
	"weikang/Data-collection/srv/product_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
)

func FlashSale(in *product.FlashSaleRequest) (*product.FlashSaleResponse, error) {
	sale, err := dao_mysql.FlashSale(in)
	if err != nil {
		return nil, err
	}
	return &product.FlashSaleResponse{Success: sale.Success}, err
}
