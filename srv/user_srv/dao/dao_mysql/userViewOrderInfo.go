package dao_mysql

import (
	"errors"
	"models/model_order/model_mysql"
)

func FindOrderInfoById(id int64) (*model_mysql.OrderInfo, error) {
	oi := &model_mysql.OrderInfo{}
	info, err := oi.FindOrderInfoById(id)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	return info, nil
}
