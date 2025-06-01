package model_mysql

import (
	"common/global"
	"fmt"
	"time"
)

type BargainUser struct {
	Id              uint32    `gorm:"column:id;type:int UNSIGNED;comment:用户参与砍价表ID;primaryKey;not null;" json:"id"`                               // 用户参与砍价表ID
	Uid             uint32    `gorm:"column:uid;type:int UNSIGNED;comment:用户ID;default:NULL;" json:"uid"`                                         // 用户ID
	BargainId       uint32    `gorm:"column:bargain_id;type:int UNSIGNED;comment:砍价商品id;default:NULL;" json:"bargain_id"`                         // 砍价商品id
	BargainPriceMin float64   `gorm:"column:bargain_price_min;type:decimal(8, 2) UNSIGNED;comment:砍价的最低价;default:NULL;" json:"bargain_price_min"` // 砍价的最低价
	BargainPrice    float64   `gorm:"column:bargain_price;type:decimal(8, 2);comment:砍价金额;default:NULL;" json:"bargain_price"`                    // 砍价金额
	Price           float64   `gorm:"column:price;type:decimal(8, 2) UNSIGNED;comment:砍掉后的价格;default:NULL;" json:"price"`                         // 砍掉的价格
	Status          uint8     `gorm:"column:status;type:tinyint UNSIGNED;comment:状态 1参与中 2 活动结束参与失败 3活动结束参与成功;default:1;" json:"status"`          // 状态 1参与中 2 活动结束参与失败 3活动结束参与成功
	CreatedAt       time.Time `gorm:"column:created_at;type:datetime(3);comment:参与时间;default:CURRENT_TIMESTAMP(3);" json:"created_at"`            // 参与时间
	DeletedAt       time.Time `gorm:"column:deleted_at;type:datetime(3);comment:是否取消;default:NULL;" json:"deleted_at"`                            // 是否取消
}

func (b *BargainUser) TableName() string {
	return fmt.Sprintf("bargain_user")
}

func (b *BargainUser) CreateBargainUser() error {
	return global.DB.Create(&b).Error
}

func (b *BargainUser) UpdateBargainPrice(id int64, price float64) error {
	return global.DB.Model(&BargainUser{}).Where("id = ?", id).Update("price", price).Error
}

func (b *BargainUser) FindBargainUserId(id int) (result *BargainUser, err error) {
	err = global.DB.Where("id = ?", id).Find(&result).Error
	return result, err
}
func (b *BargainUser) UpdateBargainStatus(id int, status int) error {
	return global.DB.Model(&BargainUser{}).Where("id = ?", id).Update("status", status).Error
}
