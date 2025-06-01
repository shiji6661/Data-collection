package dao_redis

import (
	"common/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// 查询位图中值为 1 的位的数量
func GetSignCount(uid int64) (int64, error) {
	signDate := time.Now()
	todays := signDate.Format("2006-01-02")
	key := fmt.Sprintf("sign:user:%d:%s", uid, todays)
	ctx := context.Background()
	count, err := global.Rdb.BitCount(ctx, key, &redis.BitCount{
		Start: 0,
		End:   -1,
	}).Result()
	if err != nil {
		return 0, fmt.Errorf("Failed to get sign count: %w", err)
	}
	return count, nil
}
