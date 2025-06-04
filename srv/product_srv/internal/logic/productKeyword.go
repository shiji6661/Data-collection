package logic

import (
	"weikang/Data-collection/srv/product_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
)

// todo:商品关键字查询
func ProductKeyword(in *product.ProductKeywordRequest) (*product.ProductKeywordResponse, error) {
	keyword, err := dao_mysql.FindProductKeyword(in.Keyword)
	if err != nil {
		return nil, err
	}
	return &product.ProductKeywordResponse{List: keyword}, nil
}
