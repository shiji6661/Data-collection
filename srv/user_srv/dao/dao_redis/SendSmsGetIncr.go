package dao_redis

import (
	"errors"
	"models/model_user/model_redis"
)

func SendSmsGetIncr(key string) error {
	get, _ := model_redis.SendSmsGet(key).Int()
	if get >= 3 {
		model_redis.SendSmsSet(key, get)
		return errors.New("请30分钟后重试")
	}
	return nil
}
