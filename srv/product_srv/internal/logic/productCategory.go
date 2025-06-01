package logic

import (
	"fmt"
	"product_srv/dao/dao_mysql"
	"product_srv/proto_product/product"
)

// todo:商品分类浏览
func ProductCategory(in *product.ProductCategoryRequest) (*product.ProductCategoryResponse, error) {
	list, err := dao_mysql.ProductCategory(in.CateId)
	if err != nil {
		return nil, fmt.Errorf("未找到该商品")
	}
	if list == nil {
		return nil, fmt.Errorf("查找商品失败")
	}
	return &product.ProductCategoryResponse{List: list}, nil
}
