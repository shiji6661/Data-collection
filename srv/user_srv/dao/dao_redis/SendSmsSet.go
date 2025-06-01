package dao_redis

import "models/model_user/model_redis"

func SendSmsSet(key string, code int) {
	model_redis.SendSmsSet(key, code)
}
