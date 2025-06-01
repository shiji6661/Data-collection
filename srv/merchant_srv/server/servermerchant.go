package server

import (
	"context"
	"errors"
	"merchant_srv/internal/logic"
	"merchant_srv/proto_merchant/merchant"
)

type ServerMerchant struct {
	merchant.UnimplementedMerchantServer
}

// todo:商家注册
func (s ServerMerchant) MerChantRegister(ctx context.Context, in *merchant.MerChantRegisterRequest) (*merchant.MerChantRegisterResponse, error) {
	if in.UserName == "" || in.UserPassword == "" || in.MerchantPhone == "" || in.MerchantEmail == "" {
		return nil, errors.New("请输入完整信息")
	}
	register, err := logic.MerchantRegister(in)
	if err != nil {
		return nil, err
	}
	return register, nil
}

// todo:商家登录
func (s ServerMerchant) MerchantLogin(ctx context.Context, in *merchant.MerchantLoginRequest) (*merchant.MerchantLoginResponse, error) {
	register, err := logic.MerchantLogin(in)
	if err != nil {
		return nil, err
	}
	return register, nil
}

// todo:商家短信发送
func (s ServerMerchant) SendSms(ctx context.Context, in *merchant.SendSmsRequest) (*merchant.SendSmsResponse, error) {
	register, err := logic.MerchantSms(in)
	if err != nil {
		return nil, err
	}
	return register, nil
}

// todo:商家邮箱发送
func (s ServerMerchant) SendEmail(ctx context.Context, in *merchant.SendEmailRequest) (*merchant.SendEmailResponse, error) {
	register, err := logic.SendEmail(in)
	if err != nil {
		return nil, err
	}
	return register, nil
}

// TODO 解析二维码
func (s ServerMerchant) ParseCode(ctx context.Context, in *merchant.ParseCodeRequest) (*merchant.ParseCodeResponse, error) {
	code, err := logic.ParseCode(in)
	if err != nil {
		return nil, err
	}
	return code, nil
}

// TODO 按日统计店铺数据
func (s ServerMerchant) StatisticsStoreDailyData(ctx context.Context, in *merchant.StatisticsStoreDailyDataRequest) (*merchant.StatisticsStoreDailyDataResponse, error) {
	code, err := logic.StatisticsStoreDailyData(in)
	if err != nil {
		return nil, err
	}
	return code, nil
}
