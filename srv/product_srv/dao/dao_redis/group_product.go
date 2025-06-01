package dao_redis

import (
	"context"
	"errors"
	"fmt"
	"models/model_product/model_redis"
)

const (
	ProductStockPrefix = "product:product_id:%d"
)

// SyncProductToRedis 同步商品库存
func SyncProductToRedis(ctx context.Context, productId, stock int64) error {
	key := fmt.Sprintf(ProductStockPrefix, productId)
	err := model_redis.SyncToRedis(ctx, key, stock, productId)
	if err != nil {
		return errors.New("同步redis库存失败！！！")
	}
	return nil
}

// GetFromRedis 获取redis中的库存
func GetFromRedis(ctx context.Context, productId int64) (int64, error) {
	key := fmt.Sprintf(ProductStockPrefix, productId)
	stock, err := model_redis.GetStock(ctx, key)
	if err != nil {
		return 0, errors.New("获取库存失败！！！")
	}
	return stock, nil
}

// ReduceProductNum 减少redis中的库存
func ReduceProductNum(ctx context.Context, productId, num int64) error {
	key := fmt.Sprintf(ProductStockPrefix, productId)
	err := model_redis.ReduceRedis(ctx, key, num)
	if err != nil {
		return errors.New("减少库存失败！！！")
	}
	return nil
}

// ReturnProductNum 归还redis中的库存
func ReturnProductNum(ctx context.Context, productId, num int64) error {
	key := fmt.Sprintf(ProductStockPrefix, productId)
	err := model_redis.ReturnRedis(ctx, key, num)
	if err != nil {
		return errors.New("归还库存失败！！！")
	}
	return nil
}
