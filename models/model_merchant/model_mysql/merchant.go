package model_mysql

import (
	"common/global"
	"gorm.io/gorm"
)

// todo 商家表
type Merchant struct {
	gorm.Model
	UserName       string `gorm:"type:varchar(20);comment:'商家用户名'"`
	UserPassword   string `gorm:"type:varchar(100);comment:'密码'"`
	MerchantPhone  string `gorm:"type:char(11);comment:'手机号码'"`
	MerchantEmail  string `gorm:"type:varchar(50);comment:'商家邮箱'"`
	UserState      int64  `gorm:"type:bigint(1);default:1;comment:'用户状态 0异常 1正常'"`
	MerchantAvatar string `gorm:"type:varchar(255);comment:'商家头像'"`
}

func (m Merchant) TableName() string {
	return "merchant"
}

// todo:根据商家id查询
func GetMerchantIdInfo(id int64) (merchant *Merchant, err error) {
	if err = global.DB.Raw("select * from merchant where id=? limit 1", id).Scan(&merchant).Error; err != nil {
		return nil, err
	}

	return merchant, nil
}

// todo:注册
func (m *Merchant) MerchantRegister() error {
	return global.DB.Create(&m).Error
}

// todo:根据用户名查询
func (m *Merchant) FindMerchantByUserName(userName string) error {
	return global.DB.Where("user_name = ?", userName).Limit(1).Find(&m).Error
}

// todo:根据手机号查询
func (m *Merchant) FindMerchantByPhone(phone string) error {
	return global.DB.Where("merchant_phone = ?", phone).Limit(1).Find(&m).Error
}

// todo:根据手机号查询
func (m *Merchant) FindMerchantByEmail(email string) error {
	return global.DB.Where("merchant_email = ?", email).Limit(1).Find(&m).Error
}

// todo:查询商家是否存在
func (m *Merchant) FindMerchantById(merchantId int64) error {
	return global.DB.Where("id =?", merchantId).Limit(1).Find(&m).Error
}
