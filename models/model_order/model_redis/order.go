package model_redis

import (
	"common/global"
	"context"
)

// todo:从redis中获取购物车id
func GetCartId(ctx context.Context, key string) (string, error) {
	// 判断用户是否有购物车
	result, err := global.Rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

// TODO: 判断购物车中商品是否存在
func ExistsCart(ctx context.Context, key string) bool {
	return global.Rdb.Exists(ctx, key).Val() == 1
}

// TODO: 从redis中获取商品个数
func LLenFromRedis(ctx context.Context, key string) int64 {
	return global.Rdb.LLen(ctx, key).Val()
}
