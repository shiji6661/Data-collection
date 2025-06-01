package model_mysql

import (
	"common/global"
	"errors"
	"fmt"
	"time"
)

type BargainUserHelp struct {
	Id            uint32    `gorm:"column:id;type:int UNSIGNED;comment:砍价用户帮助表ID;primaryKey;not null;" json:"id"`                    // 砍价用户帮助表ID
	Uid           uint32    `gorm:"column:uid;type:int UNSIGNED;comment:帮助的用户id;default:NULL;" json:"uid"`                           // 帮助的用户id
	BargainId     uint32    `gorm:"column:bargain_id;type:int UNSIGNED;comment:砍价商品ID;default:NULL;" json:"bargain_id"`              // 砍价商品ID
	BargainUserId uint32    `gorm:"column:bargain_user_id;type:int UNSIGNED;comment:用户参与砍价表id;default:NULL;" json:"bargain_user_id"` // 用户参与砍价表id
	Price         float64   `gorm:"column:price;type:decimal(8, 2) UNSIGNED;comment:帮助砍价多少金额;default:NULL;" json:"price"`            // 帮助砍价多少金额
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime(3);comment:添加时间;default:CURRENT_TIMESTAMP(3);" json:"created_at"` // 添加时间
}

func (b *BargainUserHelp) TableName() string {
	return fmt.Sprintf("bargain_user_help")
}

func (b *BargainUserHelp) CreateBargainUserHelp() error {
	return global.DB.Create(&b).Error
}

func (b *BargainUserHelp) BargainUserCount(Id, bargainUserId int) (int64, error) {
	var count int64
	if err := global.DB.Model(&BargainUserHelp{}).
		Where("uid = ? AND bargain_user_id = ?", Id, bargainUserId).
		Count(&count).Error; err != nil {
		return 0, errors.New("校验帮砍记录失败")
	}
	if count > 0 {
		return 0, errors.New("你已帮过该砍价活动")
	}
	return count, nil
}
