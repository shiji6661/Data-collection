package logic

import (
	"errors"
	"go.uber.org/zap"
	"merchant_srv/dao/dao_mysql"
	"merchant_srv/proto_merchant/merchant"
)

// 商家注册
func MerchantRegister(in *merchant.MerChantRegisterRequest) (*merchant.MerChantRegisterResponse, error) {
	mer, err := dao_mysql.FindMerchantByUsername(in.UserName)
	if err != nil {
		return nil, err
	}
	if mer.ID != 0 {
		zap.L().Info("商家已存在")
		return nil, errors.New("商家已存在")
	}

	register, err := dao_mysql.MerchantRegister(in)
	if err != nil {
		return nil, err
	}
	return &merchant.MerChantRegisterResponse{MerId: int64(register.ID)}, nil
}
