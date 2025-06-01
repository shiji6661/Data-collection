package dao_redis

import "models/model_user/model_redis"

func SendSmsIncr(key string) {
	incr := model_redis.SendSmsCount(key)
	if incr.Val() == 1 {
		model_redis.SendSmsRedisExpire(key)
	}
}
