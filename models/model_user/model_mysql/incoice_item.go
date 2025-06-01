package model_mysql

import (
	"common/global"
	"gorm.io/gorm"
)

type InvoiceItem struct {
	gorm.Model
	InvoiceId   int64   `gorm:"type:int;not null;comment:'关联发票id'"`
	OrderItemId int64   `gorm:"type:int;not null;comment:'关联订单商品id'"`
	ProductName string  `gorm:"type:varchar(255);not null;comment:'商品名称'"`
	Spec        string  `gorm:"type:varchar(100);not null;comment:'商品规格'"`
	Unit        string  `gorm:"type:varchar(50);not null;comment:'商品单位'"`
	Quantity    int64   `gorm:"type:int;not null;comment:'商品数量'"`
	Price       int64   `gorm:"type:int;not null;comment:'商品单价'"`
	Amount      float64 `gorm:"type:decimal(10,2);not null;comment:'金额'"`
	TaxRate     float64 `gorm:"type:decimal(5,2);not null;comment:'税率'"`
	TaxAmount   float64 `gorm:"type:decimal(10,2);not null;comment:'税额'"`
}

// todo:生成发票详情
func (it *InvoiceItem) CreateInvoiceItem() error {
	return global.DB.Create(&it).Error
}
