package dao_redis

import (
	"common/global"
	"context"
	"fmt"
	"models/model_order/model_redis"
	"strconv"
)

const (
	SyncProduct   = "product:product_id:%d"
	ProductCartId = "cart_id_%s:user_id_%d:product_id_%d"
	Cart          = "cart_id_%s:*"
	CartByUser    = "user:%s:cart_id"
)

//const (
//	SyncProducts  = "product:product_id:%d"
//	SyncProduct   = "Product_Stock:Product_id_%d"
//	ProductCartId = "cart_id_%s:user_id_%d:product_id_%d"
//	Cart          = "cart_id_%s:user_id_%d:*"
//)

type Product struct {
	CartId       string
	UserId       int64
	ProductId    int64
	Num          int64
	ProductName  string
	ProductPrice float64
	Check        int64
}

// todo:从redis中获取购物车id
func GetCartIdByUserId(ctx context.Context, userId string) (string, error) {
	cartId, err := model_redis.GetCartId(ctx, fmt.Sprintf(CartByUser, userId))
	return cartId, err
}

// todo: 判断购物车中商品是否存在
func IsCartProductExists(ctx context.Context, cartId string, userId, productId int64) bool {
	key := fmt.Sprintf(ProductCartId, cartId, userId, productId)
	return model_redis.ExistsCart(ctx, key)
}

// todo:获取商品库存
func GetProductsFromRedis(ctx context.Context, productId int64) int64 {
	key := fmt.Sprintf(SyncProduct, productId)
	return model_redis.LLenFromRedis(ctx, key)
}

// todo: 获取购物车中某个商品的信息
func GetCartProductInfo(ctx context.Context, cartId string, userId, productId int64) *Product {
	key := fmt.Sprintf(ProductCartId, cartId, userId, productId)
	all := global.Rdb.HGetAll(ctx, key).Val()
	cartid := all["cart_id"]
	userid, _ := strconv.Atoi(all["user_id"])
	productid, _ := strconv.Atoi(all["product_id"])
	num, _ := strconv.Atoi(all["num"])
	name := all["product_name"]
	price, _ := strconv.ParseFloat(all["product_price"], 64)
	check, _ := strconv.Atoi(all["check"])

	return &Product{
		CartId:       cartid,
		UserId:       int64(userid),
		ProductId:    int64(productid),
		Num:          int64(num),
		ProductName:  name,
		ProductPrice: price,
		Check:        int64(check),
	}
}
