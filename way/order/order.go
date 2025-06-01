package order

import (
	"errors"
	"models/model_order/model_mysql"
)

// 通过商品id查找订单
func FindOrderByProductId(productId int64) (o *model_mysql.Order, err error) {
	o = &model_mysql.Order{}
	err = o.FindOrderByProductId(productId)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	return o, nil
}

// todo:通过订单id查找订单详情
func FindOrderById(id int64) (o *model_mysql.Order, err error) {
	o = &model_mysql.Order{}
	err = o.FindOrderById(id)
	if err != nil {
		return nil, err
	}
	return o, nil
}
