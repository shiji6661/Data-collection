package logic

import (
	"errors"
	"weikang/Data-collection/srv/product_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
)

// todo:商品推荐
func ProductRecommend(in *product.ProductRecommendRequest) (*product.ProductRecommendResponse, error) {
	recommend, err := dao_mysql.FindProductRecommend()
	if err != nil {
		return nil, errors.New("商品查询失败")
	}
	return &product.ProductRecommendResponse{
		List: recommend,
	}, nil
}
