package model_redis

import (
	"common/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// todo 存储短信发送次数
func SendSmsCount(key string) *redis.IntCmd {
	incr := global.Rdb.Incr(context.Background(), key)
	return incr
}

// todo 错误锁定30分钟
func SendSmsRedisExpire(key string) {
	global.Rdb.Expire(context.Background(), key, time.Minute*30)
}

// todo 验证码存储10分钟
func SendSmsSet(key string, code int) {
	global.Rdb.Set(context.Background(), key, code, time.Minute*10)
}

// todo 获取发送的短信
func SendSmsGet(key string) *redis.StringCmd {
	get := global.Rdb.Get(context.Background(), key)
	return get
}

// todo 用户签到
func CheckIn(userId int64, dayOfYear int) (int64, error) {
	bit, err := global.Rdb.SetBit(context.Background(), fmt.Sprintf("checkin:%d", userId), int64(dayOfYear), 1).Result()
	if err != nil {
		return 0, err
	}
	return bit, nil
}

// todo 检查用户是否已经签到
func IsCheckIn(userId int64, dayOfYear int) (int64, error) {
	result, err := global.Rdb.GetBit(context.Background(), fmt.Sprintf("checkin:%d", userId), int64(dayOfYear)).Result()
	if err != nil {
		return 0, err
	}
	return result, nil

}

// todo 获取用户连续签到天数
func GetContinuousCheckInDays(userId int64, i int) (int64, error) {
	result, err := global.Rdb.GetBit(context.Background(), fmt.Sprintf("checkin:%d", userId), int64(i)).Result()
	if err != nil {
		return 0, err
	}
	return result, nil

}
