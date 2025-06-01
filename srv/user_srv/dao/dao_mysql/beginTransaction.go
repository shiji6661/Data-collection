package dao_mysql

import "gorm.io/gorm"

func BeginTransaction() *gorm.DB {
	transaction := U.BeginTransaction()
	return transaction
}
