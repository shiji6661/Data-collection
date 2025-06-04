package logic

import (
	"Data-collection/way/user"
	"context"
	"errors"
	"weikang/Data-collection/srv/product_srv/dao/dao_redis"
	"weikang/Data-collection/srv/product_srv/proto_product/product"

	"strconv"
)

// TODO: 购物车中商品总价
func CartProductTotalPrice(in *product.CartProductTotalPriceRequest) (*product.CartProductTotalPriceResponse, error) {
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
		return nil, err
	}
	price := dao_redis.CartProductTotalPrice(context.Background(), cartId, in.UserId)
	return &product.CartProductTotalPriceResponse{TotalPrice: float32(price)}, nil
}
