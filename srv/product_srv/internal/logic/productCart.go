package logic

import (

	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"models/model_product/model_mysql"
	"product_srv/dao/dao_redis"
	"product_srv/proto_product/product"
	"strconv"
	"way/user"
)

// TODO: 商品加入购物车
func AddCart(in *product.AddCartRequest) (*product.AddCartResponse, error) {
	// 判断用户是否存在
	info, err := user.CommonGetUserIdInfo(in.UserId)
	if err != nil {
		return nil, err
	}
	if info.ID == 0 {
		zap.L().Error("用户不存在！请重新登录！")
		return nil, errors.New("用户不存在！请重新登录！")
	}

	// 获取用户对应的购物车 ID，这里假设存在一个函数根据用户 ID 获取购物车 ID
	//user 拼接userId
	userId := strconv.Itoa(int(in.UserId))
	cartId, err := dao_redis.GetCartIdByUserId(context.Background(), userId)
	if err != nil {
		zap.L().Error("获取购物车ID失败！")
		return nil, err
	}

	// 查询商品是否存在
	p, err := model_mysql.FindProductById(in.ProductId)
	if err != nil {
		return nil, err
	}
	if p.ID == 0 {
		zap.L().Error("商品不存在！请重新选择！")
		return nil, errors.New("商品不存在！请重新选择！")
	}

	// 使用分布式锁防止并发问题
	lockKey := fmt.Sprintf("cart_lock:%d:%d", in.UserId, in.ProductId)
	acquired, err := dao_redis.AcquireLock(context.Background(), lockKey, 10) // 10秒超时
	if err != nil {
		zap.L().Error("获取分布式锁失败", zap.Error(err))
		return nil, errors.New("系统繁忙，请稍后再试")
	}
	defer dao_redis.ReleaseLock(context.Background(), lockKey)
	if !acquired {
		zap.L().Warn("获取分布式锁超时", zap.String("lockKey", lockKey))
		return nil, errors.New("系统繁忙，请稍后再试")
	}

	//判断商品库存是否足够
	stock := dao_redis.GetProductsFromRedis(context.Background(), in.ProductId)
	if stock < in.Num {
		zap.L().Error("商品库存不足！")
		return nil, errors.New("商品库存不足")
	}

	// 商品加入购物车
	// 判断购物车是否存在
	productExists := dao_redis.IsCartProductExists(context.Background(), cartId, in.UserId, in.ProductId)

	if productExists == true {
		if err = dao_redis.HIncrByProductCart(context.Background(), cartId, in.UserId, in.ProductId, in.Num); err != nil {
			zap.L().Error("购物车中商品数量自增失败！")
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
			zap.L().Error("商品加入购物车失败！")
			return nil, errors.New("商品加入购物车失败")
		}
	}

	// 商品库存减少
	if err = dao_redis.ReduceProductRedis(context.Background(), in.ProductId, int(in.Num)); err != nil {
		zap.L().Error("商品库存减少失败！")
		return nil, errors.New("商品库存减少失败")
	}

	return &product.AddCartResponse{Success: "商品添加至购物车成功！"}, nil
}
