package handler

import (
	"Api/client"
	"context"
	"merchant_srv/proto_merchant/merchant"
)

// todo:商家注册
func MerchantRegister(ctx context.Context, i *merchant.MerChantRegisterRequest) (*merchant.MerChantRegisterResponse, error) {
	merchantClient, err := client.MerchantClient(ctx, func(ctx context.Context, in merchant.MerchantClient) (interface{}, error) {
		register, err := in.MerChantRegister(ctx, i)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return merchantClient.(*merchant.MerChantRegisterResponse), nil
}

// todo:商家登录
func MerchantLogin(ctx context.Context, i *merchant.MerchantLoginRequest) (*merchant.MerchantLoginResponse, error) {
	merchantClient, err := client.MerchantClient(ctx, func(ctx context.Context, in merchant.MerchantClient) (interface{}, error) {
		register, err := in.MerchantLogin(ctx, i)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return merchantClient.(*merchant.MerchantLoginResponse), nil
}

// todo:商家短信发送
func MerchantSms(ctx context.Context, i *merchant.SendSmsRequest) (*merchant.SendSmsResponse, error) {
	merchantClient, err := client.MerchantClient(ctx, func(ctx context.Context, in merchant.MerchantClient) (interface{}, error) {
		register, err := in.SendSms(ctx, i)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return merchantClient.(*merchant.SendSmsResponse), nil
}

// todo:商家邮箱发送
func MerchantEmail(ctx context.Context, i *merchant.SendEmailRequest) (*merchant.SendEmailResponse, error) {
	merchantClient, err := client.MerchantClient(ctx, func(ctx context.Context, in merchant.MerchantClient) (interface{}, error) {
		register, err := in.SendEmail(ctx, i)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return merchantClient.(*merchant.SendEmailResponse), nil
}

// 核销二维码
func MerchantParse(ctx context.Context, i *merchant.ParseCodeRequest) (*merchant.ParseCodeResponse, error) {
	merchantClient, err := client.MerchantClient(ctx, func(ctx context.Context, in merchant.MerchantClient) (interface{}, error) {
		register, err := in.ParseCode(ctx, i)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return merchantClient.(*merchant.ParseCodeResponse), nil
}

// 按日统计店铺数据
func StatisticsStoreDailyData(ctx context.Context, i *merchant.StatisticsStoreDailyDataRequest) (*merchant.StatisticsStoreDailyDataResponse, error) {
	merchantClient, err := client.MerchantClient(ctx, func(ctx context.Context, in merchant.MerchantClient) (interface{}, error) {
		register, err := in.StatisticsStoreDailyData(ctx, i)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return merchantClient.(*merchant.StatisticsStoreDailyDataResponse), nil
}
