package model_mysql

import (
	"common/global"
	"time"
)

type Real struct {
	Id        int64     `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	Uid       int64     `gorm:"column:uid;type:int;comment:用户id;default:NULL;" json:"uid"`                      // 用户id
	RealName  string    `gorm:"column:real_name;type:varchar(10);comment:用户真实名;default:NULL;" json:"real_name"` // 用户真实名
	CardNo    string    `gorm:"column:card_no;type:char(18);comment:用户身份证号码;default:NULL;" json:"card_no"`      // 用户身份证号码
}

func (r *Real) FindRealName(card string) error {
	err := global.DB.Debug().Where("card_no = ?", card).Limit(1).Find(&r).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Real) CreateRealName() error {
	err := global.DB.Debug().Create(&r).Error
	if err != nil {
		return err
	}
	return nil
}
