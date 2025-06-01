package logic

import (
	"context"
	"errors"
	"math/rand"
	"merchant_srv/dao/dao_redis"
	"merchant_srv/proto_merchant/merchant"
	"strconv"
)

// todo:商家短信发送
func MerchantSms(in *merchant.SendSmsRequest) (*merchant.SendSmsResponse, error) {
	code := rand.Intn(9000) + 1000
	//sms, err := pkg.SendSms(in.Phone, strconv.Itoa(code))
	//if err != nil {
	//	return nil, err
	//}
	//if *sms.Body.Code != "OK" {
	//	return nil, err
	//}
	err := dao_redis.SetToRedis(context.Background(), in.Source, in.Phone, strconv.Itoa(code))
	if err != nil {
		return nil, errors.New("短信发送失败")
	}
	return &merchant.SendSmsResponse{Message: "短信发送成功！"}, nil
}
