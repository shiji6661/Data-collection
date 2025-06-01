package model_mysql

import (
	"gorm.io/gorm"
	"time"
)

type InvoiceDelivery struct {
	gorm.Model
	InvoiceId      int64     `gorm:"type:int;not null;comment:'关联发票Id'" json:"invoice_id"`                                // 关联发票ID
	ExpressCompany string    `gorm:"type:varchar(50);not null;comment:'快递公司'" json:"express_company"`                     // 快递公司
	ExpressNo      string    `gorm:"type:varchar(100);not null;comment:'快递单号'" json:"express_no"`                         // 快递单号
	DeliveryStatus int64     `gorm:"type:int;not null;default:1;comment:'寄送状态：1-待发货，2-已发货，3-已签收'" json:"delivery_status"` // 寄送状态：1-待发货，2-已发货，3-已签收
	DeliveryTime   time.Time `gorm:"type:datetime;not null;comment:'发货时间'" json:"delivery_time"`                          // 发货时间
	ReceivedTime   time.Time `gorm:"type:datetime;not null;comment:'签收时间'" json:"received_time"`                          // 签收时间
	Address        string    `gorm:"type:varchar(255);not null;comment:'收货地址'" json:"address"`                            // 收件地址
	Contact        string    `gorm:"type:varchar(50);not null;comment:'联系人'" json:"contact"`                              // 联系人
	Phone          string    `gorm:"type:char(11);not null;comment:'联系电话'" json:"phone"`                                  // 联系电话
}
