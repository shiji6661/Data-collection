package model_mysql

import (
	"common/global"
	"errors"
	"gorm.io/gorm"
	"time"
)

type UserSign struct {
	Id       int32     `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	Uid      int32     `gorm:"column:uid;type:int;comment:用户uid;not null;default:0;" json:"uid"`      // 用户uid
	Title    string    `gorm:"column:title;type:varchar(255);comment:签到说明;not null;" json:"title"`    // 签到说明
	Number   int32     `gorm:"column:number;type:int;comment:获得积分;not null;default:0;" json:"number"` // 获得积分
	AddTime  time.Time `gorm:"column:add_time;type:datetime;comment:添加时间;not null;" json:"add_time"`  // 添加时间
	IsMakeup int32     `gorm:"column:is_makeup;type:int;comment:是否签到;default:0;" json:"is_makeup"`    // 是否补签
}

func (u *UserSign) TableName() string {
	return "user_sign"
}

func (u *UserSign) UserSigns(point int, frequency int64) error {
	tx := global.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 插入 user_sign 记录
	if err := tx.Create(u).Error; err != nil {
		tx.Rollback()
		return err
	}

	//更新用户签到积分
	if err := tx.Debug().Model(&User{}).Where("id = ?", u.Uid).Update("user_point", gorm.Expr("user_point + ?", point)).Error; err != nil {
		tx.Rollback()
		return err
	}
	//更新用户连续签到记录
	if err := tx.Debug().Model(&User{}).Where("id = ?", u.Uid).Update("consecutive_sign_ins", gorm.Expr("consecutive_sign_ins + ?", frequency)).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (u *UserSign) IsUserSignIn(uid int64, today time.Time, tomorrow time.Time) error {
	return global.DB.Where("uid = ? AND add_time >= ? AND add_time < ?", uid, today, tomorrow).First(&UserSign{}).Error

}

func (u *UserSign) ResetSignInDays(id int64) error {
	return global.DB.Model(&User{}).Where("id = ?", id).Update("consecutive_sign_ins", 0).Error
}

func GetLastSignDate(uid int64) (time.Time, error) {
	var record UserSign
	result := global.DB.Where("uid = ?", uid).Order("add_time DESC").First(&record)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return time.Time{}, nil
		}
		return time.Time{}, result.Error
	}
	return record.AddTime, nil
}
