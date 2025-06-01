package logic

import (
	"Api/pkg"
	"common/utils"
	"context"
	"errors"
	"go.uber.org/zap"
	"merchant_srv/dao/dao_mysql"
	"merchant_srv/dao/dao_redis"
	"merchant_srv/proto_merchant/merchant"
)

// 商家登录一体化
func MerchantLogin(in *merchant.MerchantLoginRequest) (*merchant.MerchantLoginResponse, error) {
	switch in.LoginType {
	case merchant.MerchantLoginType_MERCHANT_LOGIN_TYPE_USERNAME_PASSWORD:
		if in.UserName == "" || in.UserPassword == "" {
			return nil, errors.New("用户名或密码不能为空")
		}
		username, err := dao_mysql.FindMerchantByUsername(in.UserName)
		if err != nil {
			return nil, err
		}
		if username.ID == 0 {
			zap.L().Info("商家不存在！请重新登录")
			return nil, errors.New("商家不存在！请重新登录")
		}
		if username.UserPassword != utils.Sha256Encrypt(in.UserPassword) {
			zap.L().Error("用户名或密码错误")
			return nil, errors.New("用户名或密码错误")
		}
		claims := pkg.CustomClaims{ID: username.ID}
		token, err := pkg.NewJWT("Merchant").CreateToken(claims)
		if err != nil {
			return nil, err
		}
		return &merchant.MerchantLoginResponse{
			Token: token,
		}, nil
	case merchant.MerchantLoginType_MERCHANT_LOGIN_TYPE_PHONE_PASSWORD:
		if in.MerchantPhone == "" || in.UserPassword == "" {
			return nil, errors.New("手机号或密码不能为空")
		}
		phone, err := dao_mysql.FindMerchantByPhone(in.MerchantPhone)
		if err != nil {
			return nil, err
		}
		if phone.ID == 0 {
			zap.L().Info("手机号码不存在！请重新登录")
			return nil, errors.New("手机号码不存在！请重新登录")
		}
		if phone.UserPassword != utils.Sha256Encrypt(in.UserPassword) {
			zap.L().Error("手机号码或密码错误")
			return nil, errors.New("手机号码或密码错误")
		}
		// 验证验证码
		redis, err := dao_redis.GetLoginCodeFromRedis(context.Background(), in.MerchantPhone)
		if err != nil {
			return nil, err
		}
		if redis != in.Code {
			zap.L().Info("验证码错误！")
			return nil, errors.New("验证码错误！")
		}
		claims := pkg.CustomClaims{ID: phone.ID}
		token, err := pkg.NewJWT("Merchant").CreateToken(claims)
		if err != nil {
			return nil, err
		}
		return &merchant.MerchantLoginResponse{
			Token: token,
		}, nil
	case merchant.MerchantLoginType_MERCHANT_LOGIN_TYPE_EMAIL_PASSWORD:
		if in.MerchantEmail == "" || in.UserPassword == "" {
			return nil, errors.New("邮箱或密码不能为空")
		}
		email, err := dao_mysql.FindMerchantByEmail(in.MerchantEmail)
		if err != nil {
			return nil, err
		}
		if email.ID == 0 {
			zap.L().Info("邮箱不存在！请重新登录")
			return nil, errors.New("邮箱不存在！请重新登录")
		}

		if email.UserPassword != utils.Sha256Encrypt(in.UserPassword) {
			zap.L().Error("邮箱或密码错误")
			return nil, errors.New("邮箱或密码错误")
		}
		redis, err := dao_redis.GetEmailCodeFromRedis(context.Background(), in.MerchantEmail)
		if err != nil {
			return nil, err
		}
		if redis != in.Code {
			zap.L().Info("验证码错误！")
			return nil, errors.New("验证码错误！")
		}
		claims := pkg.CustomClaims{ID: email.ID}
		token, err := pkg.NewJWT("Merchant").CreateToken(claims)
		if err != nil {
			return nil, err
		}
		return &merchant.MerchantLoginResponse{
			Token: token,
		}, nil
	case merchant.MerchantLoginType_MERCHANT_LOGIN_TYPE_UNSPECIFIED:
		return nil, errors.New("登录类型错误")
	default:
		return nil, errors.New("登录类型错误")
	}
}
