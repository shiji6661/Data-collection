package dao_mysql

import "models/model_order/model_mysql"

func UpdateProductPay(pay int64, proId int64) error {
	o := &model_mysql.Order{}
	err := o.UpdateProductPay(pay, proId)
	if err != nil {
		return err
	}
	return nil
}
