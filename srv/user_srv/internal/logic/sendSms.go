package logic

import (
	user2 "common/utils/user"
	"errors"
	"math/rand"
	"weikang/Data-collection/srv/user_srv/dao/dao_redis"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
)

// todo:短信验证码逻辑
func SendSms(in *user.SendSmsRequest) (*user.SendSmsResponse, error) {
	if len(in.Mobile) != 11 {
		return nil, errors.New("手机号长度不对")
	}
	if !user2.ValidateMobile(in.Mobile) {
		return nil, errors.New("手机号格式不对")
	}
	err := dao_redis.SendSmsGetIncr("SendSmsIncr" + in.Source + in.Mobile)
	if err != nil {
		return nil, err
	}
	code := rand.Intn(9000) + 1000
	/*_, err := pkg.SendSms(in.UserPhone, strconv.Itoa(code))
	if err != nil {
		return nil, err
	}*/
	dao_redis.SendSmsSet("SendSms"+in.Source+in.Mobile, code)
	dao_redis.SendSmsIncr("SendSmsIncr" + in.Source + in.Mobile)

	return &user.SendSmsResponse{Greet: "发送成功"}, nil
}
