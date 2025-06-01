package dao_mysql

import (
	"common/utils"
	merchant2 "common/utils/merchant"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"merchant_srv/proto_merchant/merchant"
	"models/model_merchant/model_mysql"
)

const MerStatus = 1 //用户状态正常

// MerchantRegister 商户注册
func MerchantRegister(in *merchant.MerChantRegisterRequest) (*model_mysql.Merchant, error) {
	err := merchant2.Phone(in.MerchantPhone)
	if err != nil {
		return nil, err
	}
	if !merchant2.Email(in.MerchantEmail) {
		return nil, err
	}

	m := &model_mysql.Merchant{
		Model:          gorm.Model{},
		UserName:       in.UserName,
		UserPassword:   utils.Sha256Encrypt(in.UserPassword),
		MerchantPhone:  in.MerchantPhone,
		MerchantEmail:  in.MerchantEmail,
		UserState:      MerStatus,
		MerchantAvatar: in.MerchantAvatar,
	}
	err = m.MerchantRegister()
	if err != nil {
		zap.L().Info("商户注册失败")
		return nil, err
	}
	return m, nil
}

// 查找Merchant商户名称是否存在
func FindMerchantByUsername(username string) (m *model_mysql.Merchant, err error) {
	m = &model_mysql.Merchant{}
	err = m.FindMerchantByUserName(username)
	if err != nil {
		zap.L().Error("")
		return nil, err
	}
	return m, nil
}

// todo:通过手机号查找
func FindMerchantByPhone(phone string) (m *model_mysql.Merchant, err error) {
	m = &model_mysql.Merchant{}
	err = m.FindMerchantByPhone(phone)
	if err != nil {
		zap.L().Error("")
		return nil, err
	}
	return m, nil
}

// todo:通过邮箱
func FindMerchantByEmail(email string) (m *model_mysql.Merchant, err error) {
	m = &model_mysql.Merchant{}
	err = m.FindMerchantByEmail(email)
	if err != nil {
		zap.L().Error("")
		return nil, err
	}
	return m, nil
}
