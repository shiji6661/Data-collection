package dao_redis

import (
	"common/global"
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"models/model_product/model_redis"
	"strconv"
)

const (
	SyncProducts  = "product:product_id:%d"
	SyncProduct   = "Product_Stock:Product_id_%d"
	ProductCartId = "cart_id_%s:user_id_%d:product_id_%d"
	Cart          = "cart_id_%s:user_id_%d:*"
)

// todo:同步商品到redis中
func SyncProductsToRedis(ctx context.Context, productId, num int) error {
	key := fmt.Sprintf(SyncProduct, productId)
	err := model_redis.LPushToRedis(ctx, key, num, productId)
	if err != nil {
		return errors.New("同步商品到redis失败")
	}
	return nil
}

// todo: 根据用户id获取购物车id
func GetCartIdByUserId(ctx context.Context, userId string) (string, error) {
	cartId, err := global.Rdb.Get(ctx, fmt.Sprintf("user:%s:cart_id", userId)).Result()
	if err == redis.Nil {
		// 如果购物车 ID 不存在，生成一个新的
		newCartId := uuid.New().String()
		// 将新的购物车 ID 存储到 Redis 中
		err = global.Rdb.Set(ctx, fmt.Sprintf("user:%s:cart_id", userId), newCartId, 0).Err()
		if err != nil {
			return "", err
		}
		return newCartId, nil
	} else if err != nil {
		return "", err
	}
	return cartId, nil
}

func IsCartProductExists(ctx context.Context, cartId string, userId int64, productId int64) bool {
	key := fmt.Sprintf(ProductCartId, cartId, userId, productId)
	return global.Rdb.Exists(ctx, key).Val() == 1

}

func HIncrByProductCart(ctx context.Context, cartid string, id int64, id2 int64, num int64) error {
	key := fmt.Sprintf(ProductCartId, cartid, id, id2)

	err := model_redis.HIncrCart(ctx, key, "num", int(num))
	if err != nil {
		return errors.New("商品自增失败")
	}
	return nil
}

// todo: 减少redis中商品数量
func ReduceProductRedis(ctx context.Context, productId int64, num int) error {
	key := fmt.Sprintf(SyncProducts, productId)
	err := model_redis.ReduceProduct(ctx, key, num)
	if err != nil {
		return err
	}
	return nil
}

// todo:添加商品到购物车
func AddProductToCart(ctx context.Context, cart map[string]interface{}) error {
	key := fmt.Sprintf(ProductCartId, cart["cart_id"], cart["user_id"], cart["product_id"])
	err := model_redis.HMSetToRedis(ctx, key, cart)
	if err != nil {
		return errors.New("添加商品到购物车失败")
	}
	return nil
}

// todo:获取商品库存
func GetProductsFromRedis(ctx context.Context, productId int64) int64 {
	key := fmt.Sprintf(SyncProducts, productId)
	return model_redis.LLenFromRedis(ctx, key)
}

// todo: 统计购物车中商品的价格
func CartProductTotalPrice(ctx context.Context, cartId string, userId int64) float64 {
	var totalPrice float64
	key := fmt.Sprintf(Cart, cartId, userId)
	// 获取购物车中的商品
	keys := model_redis.KeysCart(ctx, key)
	for _, s := range keys {
		// 获取购物车中的商品
		all := model_redis.HGetAll(ctx, s)
		num, _ := strconv.Atoi(all["num"])
		price, _ := strconv.ParseFloat(all["product_price"], 64)
		check, _ := strconv.Atoi(all["check"])
		if check == 1 && num != 0 {
			totalPrice += float64(num) * price
		}
	}
	return totalPrice
}
