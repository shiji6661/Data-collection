package dao_redis

import (
	"common/global"
	"context"
	"errors"
	"fmt"
	"time"
)

const (
	SmsCode       = "sendSms"
	SmsLogin      = "sendSmslogin"
	SmsRegister   = "sendSmsregister"
	SendEmailCode = "sendEmailCode"
)

// 存入redis验证码
func SetToRedis(ctx context.Context, source, phone string, code string) error {
	key := SmsCode + source + phone
	err := global.Rdb.Set(ctx, key, code, time.Minute*5).Err()
	if err != nil {
		return err
	}
	return nil
}

// 从redis中获取验证码
func GetLoginCodeFromRedis(ctx context.Context, phone string) (string, error) {
	key := SmsLogin + phone
	result := global.Rdb.Get(ctx, key).Val()
	if result == "" {
		return "", errors.New("验证码不存在")
	}
	fmt.Println(result)
	return result, nil
}

// 发送邮件验证码到redis
func SetEmailCodeToRedis(ctx context.Context, email string, code string) error {
	key := SendEmailCode + email
	err := global.Rdb.Set(ctx, key, code, time.Minute*5).Err()
	if err != nil {
		return err
	}
	return nil
}

// 从redis中获取email验证码
func GetEmailCodeFromRedis(ctx context.Context, email string) (string, error) {
	key := SendEmailCode + email
	result := global.Rdb.Get(ctx, key).Val()
	if result == "" {
		return "", errors.New("验证码不存在")
	}
	return result, nil
}
