package dao_redis

import (
	"errors"
	"models/model_user/model_redis"
)

func SendSmsGet(key string, code string) error {
	get := model_redis.SendSmsGet(key)
	if get.Val() != code {
		return errors.New("验证码错误")
	}
	return nil
}
