package logic

import (
	pkg "common/pkg/pkg_merchant"
	"context"
	"errors"
	"math/rand"
	"merchant_srv/dao/dao_redis"
	"merchant_srv/proto_merchant/merchant"
	"strconv"
)

// 邮箱发送
func SendEmail(in *merchant.SendEmailRequest) (*merchant.SendEmailResponse, error) {
	code := rand.Intn(9000) + 1000
	pkg.SendEmail(in.Email, strconv.Itoa(code))
	err := dao_redis.SetEmailCodeToRedis(context.Background(), in.Email, strconv.Itoa(code))
	if err != nil {
		return nil, errors.New("存入redis失败")
	}
	return &merchant.SendEmailResponse{Message: "邮箱发送成功！"}, nil
}
