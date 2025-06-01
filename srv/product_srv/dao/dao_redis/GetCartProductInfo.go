package dao_redis

import (
	"common/global"
	"context"
	"fmt"
	"strconv"
)

type Product struct {
	CartId       string
	UserId       int64
	ProductId    int64
	Num          int64
	ProductName  string
	ProductPrice float64
	Check        int64
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
