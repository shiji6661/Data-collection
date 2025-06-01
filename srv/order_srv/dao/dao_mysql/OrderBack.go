package dao_mysql

import "models/model_order/model_mysql"

func OrderBack(id string, status int64) error {
	order := &model_mysql.Order{}
	err := order.FindOrderOrderSn(id, status)
	if err != nil {
		return err
	}
	return nil
}
