package logic

import (
	"common/global"
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"product_srv/dao/dao_redis"
	"product_srv/proto_product/product"

	"strconv"
	pro "way/product"
	"way/user"
)

func AddToCart(in *product.AddToCartRequest) (*product.AddToCartResponse, error) {

	ctx := context.Background()

	//购物车商品数量上限限制
	if in.Num > global.Shopping_cart_quantity_limit {
		return nil, errors.New("商品添加数量超过上限")
	}

	//查找用户是否存在
	users, err := user.CommonGetUserIdInfo(in.UserId)
	if err != nil {
		return nil, errors.New("用户信息查询失败")
	}
	if users == nil {
		return nil, errors.New("用户信息为空")
	}

	//user 拼接userId
	userId := strconv.Itoa(int(in.UserId))
	cartId, err := dao_redis.GetCartIdByUserId(context.Background(), userId)
	if err != nil {
		zap.L().Error("获取购物车ID失败！")
		return nil, err
	}

	fmt.Println(cartId)

	//查找商品
	p, err := pro.GetCommonFindProductById(in.ProductId)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, errors.New("商品信息为空")
	}

	//判断商品库存是否足够
	stock := dao_redis.GetProductsFromRedis(context.Background(), in.ProductId)
	if stock < in.Num {
		zap.L().Error("商品库存不足！")
		return nil, errors.New("商品库存不足")
	}
	fmt.Println(stock)
	//判断购物车是否存在
	exists := dao_redis.IsCartProductExists(ctx, cartId, in.UserId, in.ProductId)
	if exists == true {
		err = dao_redis.HIncrByProductCart(ctx, cartId, in.UserId, in.ProductId, in.Num)
		if err != nil {
			return nil, errors.New("购物车中商品数量自增失败！")
		}
	} else {
		cart := map[string]interface{}{
			"cart_id":       cartId,
			"user_id":       in.UserId,
			"product_id":    in.ProductId,
			"num":           in.Num,
			"product_name":  p.StoreName,
			"product_price": p.Price * float64(in.Num),
			"check":         1,
		}
		err = dao_redis.AddProductToCart(context.Background(), cart)
		if err != nil {
			return nil, errors.New("商品加入购物车失败！")
		}
	}

	fmt.Println(p.StoreName)
	fmt.Println(exists)

	if err = dao_redis.ReduceProductRedis(context.Background(), in.ProductId, int(in.Num)); err != nil {
		zap.L().Error("商品库存减少失败！")
		return nil, errors.New("商品库存减少失败")
	}

	return &product.AddToCartResponse{Success: true}, nil
}
