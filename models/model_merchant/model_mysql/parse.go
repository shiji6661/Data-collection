package model_mysql

import (
	"common/global"
	"time"
)

type MerchantParse struct {
	Id        uint32    `gorm:"column:id;type:int UNSIGNED;comment:商家解析记录表;primaryKey;not null;" json:"id"`      // 商家解析记录表
	UserId    int32     `gorm:"column:user_id;type:int;comment:用户id;not null;" json:"user_id"`                   // 用户id
	OrderId   int32     `gorm:"column:order_id;type:int;comment:订单id;not null;" json:"order_id"`                 // 订单id
	MerId     int32     `gorm:"column:mer_id;type:int;comment:商家id;not null;" json:"mer_id"`                     // 商家id
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(3);comment:创建时间;not null;" json:"created_at"`     // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(3);comment:修改时间;not null;" json:"updated_at"`     // 修改时间
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime(3);comment:删除时间;default:NULL;" json:"deleted_at"` // 删除时间
}

func (m *MerchantParse) TableName() string {
	return "merchant_parse"
}

func (m *MerchantParse) CreateParse() error {
	return global.DB.Create(&m).Error
}
