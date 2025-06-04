package model_mysql

import (
	"common/global"
	"common/utils"
	"time"
	"user_srv/proto_user/user"
)

type Address struct {
	Id                  uint32    `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	CreatedAt           time.Time `gorm:"column:created_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt           time.Time `gorm:"column:deleted_at;type:datetime(3);comment:删除 ;default:NULL;" json:"deleted_at"`                       // 删除
	UserId              uint32    `gorm:"column:user_id;type:int UNSIGNED;comment:用户id;default:NULL;" json:"user_id"`                           // 用户id
	UserName            string    `gorm:"column:user_name;type:varchar(255);comment:用户收货昵称;not null;" json:"user_name"`                         // 用户收货昵称
	UserPhone           string    `gorm:"column:user_phone;type:char(11);comment:用户手机号;not null;" json:"user_phone"`                            // 用户手机号
	UserProvince        string    `gorm:"column:user_province;type:varchar(255);comment:用户地址省份;not null;" json:"user_province"`                 // 用户地址省份
	UserCity            string    `gorm:"column:user_city;type:varchar(255);comment:用户地址城市;not null;" json:"user_city"`                         // 用户地址城市
	UserCounty          string    `gorm:"column:user_county;type:varchar(255);comment:用户地址市区;not null;" json:"user_county"`                     // 用户地址市区
	UserDetailedAddress string    `gorm:"column:user_detailed_address;type:varchar(255);comment:用户详细地址;not null;" json:"user_detailed_address"` // 用户详细地址
}

func (a *Address) TableName() string {
	return "address"
}

func (a *Address) UpdateUser(in *user.UserModifyRequest) (add *Address, err error) {
	tx := global.DB.Begin()
	add = &Address{
		UserName:            in.UserName,
		UserPhone:           in.UserPhone,
		UserProvince:        in.UserProvince,
		UserCity:            in.UserCity,
		UserCounty:          in.UserCounty,
		UserDetailedAddress: in.UserDetailedAddress,
	}
	err = tx.Debug().Model(&Address{}).Where("user_id = ?", in.Id).Updates(&add).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	users := &User{
		UserName:     in.UserName,
		UserPassword: utils.Md5(in.UserPassword),
		UserPhone:    in.UserPhone,
		UserEmail:    in.UserEmail,
	}
	err = tx.Debug().Model(&User{}).Where("id = ?", in.Id).Updates(&users).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return add, nil
}
