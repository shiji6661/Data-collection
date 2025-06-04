package dao_redis

import (
	"common/global"
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"models/model_product/model_redis"
	"product_srv/proto_product/product"

	"time"

	"strconv"
)

const (
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
	key := fmt.Sprintf(SyncProduct, productId)
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
	key := fmt.Sprintf(SyncProduct, productId)
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

// todo:购物车中商品列表
func ProductCartList(ctx context.Context, cartId string, userId int64) []*product.ProductItems {
	key := fmt.Sprintf(Cart, cartId, userId)
	// 获取购物车中的商品
	keys := model_redis.KeysCart(ctx, key)
	var productList []*product.ProductItems
	for _, s := range keys {
		// 获取购物车中的商品
		all := model_redis.HGetAll(ctx, s)
		// 转换为商品列表
		userid, _ := strconv.Atoi(all["user_id"])
		productId, _ := strconv.Atoi(all["product_id"])
		num, _ := strconv.Atoi(all["num"])
		check, _ := strconv.Atoi(all["check"])
		price, _ := strconv.ParseFloat(all["product_price"], 64)
		productList = append(productList, &product.ProductItems{
			UserId:      int64(userid),
			ProductId:   int64(productId),
			Num:         int64(num),
			Price:       float32(price),
			ProductName: all["product_name"],
			Check:       int64(check),
		})
	}
	return productList
}

// todo: 统计购物车中商品的数量
func CartProductCount(ctx context.Context, cartId string, userId int64) int64 {
	var count int64
	key := fmt.Sprintf(Cart, cartId, userId)
	// 获取购物车中的商品
	keys := model_redis.KeysCart(ctx, key)
	for _, s := range keys {
		// 获取购物车中的商品
		all := model_redis.HGetAll(ctx, s)
		userid, _ := strconv.Atoi(all["user_id"])
		num, _ := strconv.Atoi(all["num"])
		check, _ := strconv.Atoi(all["check"])
		// 统计购物车中商品的数量
		if check == 1 && userid == int(userId) {
			count += int64(num)
		}

	}
	return count
}

// todo: 清空购物车
func ClearProductCart(ctx context.Context, cartId string, userId int64) error {
	key := fmt.Sprintf(Cart, cartId, userId)
	// 获取购物车中的商品
	keys := model_redis.KeysCart(ctx, key)
	for _, s := range keys {
		err := model_redis.RemoveCart(ctx, s)
		if err != nil {
			return errors.New("清空购物车失败！")
		}
	}
	return nil
}

// todo: 删除购物车中的商品
func DeleteProductFromCart(ctx context.Context, cartId string, userId, productId int64) error {
	key := fmt.Sprintf(ProductCartId, cartId, userId, productId)
	err := model_redis.RemoveCart(ctx, key)
	if err != nil {
		return errors.New("移除购物车中商品失败！")
	}
	return nil
}

// TODO: 查询商品根据商品id
func FindProductById(productId int64) (p *Product, err error) {
	err = global.DB.Raw("select * from product where id=? limit 1", productId).Scan(&p).Error
	if err != nil {
		return nil, err
	}
	return p, nil

}

// todo: 直接修改购物车中的商品数量
func UpdateProductCart(ctx context.Context, cartId string, userId, productId int64, num int64) error {
	key := fmt.Sprintf(ProductCartId, cartId, userId, productId)
	err := model_redis.UpdateCart(ctx, key, "num", num)
	if err != nil {
		return errors.New("修改商品数量失败！")
	}
	return nil
}

// 分布式锁相关常量
const (
	lockKeyPrefix      = "lock:"          // 锁键前缀
	defaultLockTimeout = 10 * time.Second // 默认锁超时时间
)

// AcquireLock 尝试获取分布式锁
// key: 锁的唯一标识
// timeout: 锁的过期时间，建议大于业务处理时间
// 返回: 是否获取成功, 错误信息
func AcquireLock(ctx context.Context, key string, timeout time.Duration) (bool, error) {
	// 生成唯一的客户端标识，确保释放锁时的安全性
	clientID, err := generateClientID()
	if err != nil {
		return false, err
	}

	// 构建完整的锁键
	lockKey := lockKeyPrefix + key

	// 使用SET命令实现原子性获取锁
	// NX: 只有键不存在时才设置
	// EX: 设置过期时间(秒)
	// 使用客户端ID作为值，确保释放锁时的正确性
	result, err := global.Rdb.SetNX(ctx, lockKey, clientID, timeout).Result()
	if err != nil {
		return false, err
	}

	// 返回是否获取成功
	return result, nil
}

// ReleaseLock 释放分布式锁
// key: 锁的唯一标识
// 返回: 是否释放成功, 错误信息
func ReleaseLock(ctx context.Context, key string) (bool, error) {
	// 构建完整的锁键
	lockKey := lockKeyPrefix + key

	// 获取锁的当前值（客户端ID）
	currentID, err := global.Rdb.Get(ctx, lockKey).Result()
	if err != nil {
		// 锁不存在或已过期，视为释放成功
		if err == redis.Nil {
			return true, nil
		}
		return false, err
	}

	// 使用Lua脚本确保释放锁的原子性
	// 脚本逻辑: 仅当锁的值等于当前客户端ID时才删除锁
	script := `
		if redis.call("get", KEYS[1]) == ARGV[1] then
			return redis.call("del", KEYS[1])
		else
			return 0
		end
	`

	// 执行Lua脚本
	result, err := global.Rdb.Eval(ctx, script, []string{lockKey}, currentID).Result()
	if err != nil {
		return false, err
	}

	// 转换结果为布尔值
	return result.(int64) == 1, nil
}

// generateClientID 生成唯一的客户端ID
func generateClientID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// TryAcquireLockWithRetry 带重试机制的获取锁
func TryAcquireLockWithRetry(ctx context.Context, key string, timeout time.Duration, retryTimes int, retryInterval time.Duration) (bool, error) {
	for i := 0; i < retryTimes; i++ {
		acquired, err := AcquireLock(ctx, key, timeout)
		if err != nil {
			return false, err
		}
		if acquired {
			return true, nil
		}
		// 等待一段时间后重试
		time.Sleep(retryInterval)
	}
	return false, nil
}
