package model_redis

import (
	"common/global"
	"context"
	"errors"
	"fmt"
	"time"
)

// todo 同步库存
func SyncToRedis(ctx context.Context, key string, val1, val2 int64) error {
	for i := 0; i < int(val1); i++ {
		err := global.Rdb.LPush(ctx, key, val2).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

// todo 获取库存
func GetStock(ctx context.Context, key string) (int64, error) {
	result, err := global.Rdb.LLen(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}

// todo 商品增加
func HIncrCart(ctx context.Context, key string, s string, num int) error {
	for i := 0; i < num; i++ {
		err := global.Rdb.HIncrBy(ctx, key, s, 1).Err()
		if err != nil {
			return errors.New("商品自增失败")
		}
	}
	return nil
}

// todo 商品添加购物车
func AddCart(ctx context.Context, key string, cart map[string]interface{}) error {
	err := global.Rdb.HMSet(ctx, key, cart).Err()
	if err != nil {
		return errors.New("添加购物车失败")
	}
	return nil
}

// todo 商品扣减库存
func ReduceProductRedis(ctx context.Context, key string, num int) error {

	for i := 0; i < num; i++ {
		err := global.Rdb.LPop(ctx, key).Err()
		if err != nil {
			return errors.New("减少失败！")
		}
	}
	return nil
}

// 减少库存
func ReduceRedis(ctx context.Context, key string, val int64) error {
	for i := 0; i < int(val); i++ {
		err := global.Rdb.LPop(ctx, key).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

// TODO: LPush商品到redis中
func LPushToRedis(ctx context.Context, key string, val1 int, val2 int) error {
	for i := 0; i < 100; i++ {
		err := global.Rdb.LPush(ctx, key, val2).Err()
		if err != nil {
			return errors.New("存入redis失败")
		}
	}
	return nil
}

// 返回库存
func ReturnRedis(ctx context.Context, key string, val int64) error {
	for i := 0; i < int(val); i++ {
		err := global.Rdb.LPush(ctx, key, val).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

// TODO: 添加商品到redis中
func HMSetToRedis(ctx context.Context, key string, val1 map[string]interface{}) error {
	err := global.Rdb.HMSet(ctx, key, val1).Err()
	if err != nil {
		return errors.New("存入redis失败")
	}
	return nil
}

// TODO: 减少redis中商品数量
func ReduceProduct(ctx context.Context, key string, num int) error {
	for i := 0; i < num; i++ {
		err := global.Rdb.LPop(ctx, key).Err()
		if err != nil {
			return errors.New("减少失败！")
		}
	}
	return nil
}

// TODO: 从redis中获取商品个数
func LLenFromRedis(ctx context.Context, key string) int64 {
	return global.Rdb.LLen(ctx, key).Val()
}

// TODO: 获取redis中商品
func KeysCart(ctx context.Context, key string) []string {
	return global.Rdb.Keys(ctx, key).Val()
}

// TODO: 获取redis中商品
func HGetAll(ctx context.Context, key string) map[string]string {
	return global.Rdb.HGetAll(ctx, key).Val()
}

func CreateSpike(productID uint, stock int) error {
	key := fmt.Sprintf("SpikeProduct:%d", productID)
	lockKey := fmt.Sprintf("Lock:SpikeProducts:%d", productID)
	lockValue := "locked"
	result, err := global.Rdb.SetNX(context.Background(), lockKey, lockValue, time.Second*60).Result()
	if err != nil {
		return fmt.Errorf("获取锁失败 err:%v", err)
	}
	if !result {
		return fmt.Errorf("无法获取锁")
	}

	defer func() {
		if err := global.Rdb.Del(context.Background(), lockKey).Err(); err != nil {
			fmt.Printf("释放锁失败%s\n", lockKey)

		}
	}()
	for i := 0; i < stock; i++ {
		err = global.Rdb.LPush(context.Background(), key, productID).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func SpikeLen(productID int) int64 {
	key := fmt.Sprintf("SpikeProduct:%d", productID)
	return global.Rdb.LLen(context.Background(), key).Val()
}

func SpikeRPop(productID int) error {
	key := fmt.Sprintf("SpikeProduct:%d", productID)
	return global.Rdb.RPop(context.Background(), key).Err()
}

// TODO: 移除redis中商品
func RemoveCart(ctx context.Context, key string) error {
	err := global.Rdb.Del(ctx, key).Err()
	if err != nil {
		return errors.New("移除失败！")
	}
	return nil
}

// TODO: 直接修改redis中商品数量
func UpdateCart(ctx context.Context, key string, field string, num int64) error {
	err := global.Rdb.HMSet(ctx, key, field, num).Err()
	if err != nil {
		return errors.New("修改失败！")
	}
	return nil
}
