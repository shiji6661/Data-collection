package product

import (
	"errors"
	"models/model_product/model_mysql"
)

// TODO: 查询商品根据商品id
func GetCommonFindProductById(productId int64) (product *model_mysql.Product, err error) {

	var a model_mysql.Product
	pro, err := a.FindProductById(productId)

	product = &model_mysql.Product{}
	pro, err = product.FindProductById(productId)

	if err != nil {
		return nil, errors.New("商品信息查询失败")
	}
	return pro, nil
}
