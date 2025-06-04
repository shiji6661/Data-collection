package logic

import (
	"Data-collection/way/user"
	"context"
	"errors"
	"product_srv/dao/dao_redis"
	"product_srv/proto_product/product"

	"strconv"
)

// todo 购物车计算总价
func CartTotalPrice(in *product.CartTotalPriceRequest) (*product.CartTotalPriceResponse, error) {
	// 判断用户是否存在
	info, err := user.CommonGetUserIdInfo(in.UserId)
	if err != nil {
		return nil, err
	}
	if info.ID == 0 {
		return nil, errors.New("用户不存在！请重新登录！")
	}
	cartId, err := dao_redis.GetCartIdByUserId(context.Background(), strconv.Itoa(int(in.UserId)))
	if err != nil {
		return nil, errors.New("购物车ID查询失败")
	}

	price := dao_redis.CartProductTotalPrice(context.Background(), cartId, in.UserId)
	return &product.CartTotalPriceResponse{
		TotalPrice: float32(price),
	}, nil
}
