package logic

import (
	"context"
	"errors"
	"product_srv/dao/dao_redis"
	"product_srv/proto_product/product"

	"strconv"
	"way/user"
)

// TODO: 统计购物车中商品总数量
func CartProductCount(in *product.CartProductCountRequest) (*product.CartProductCountResponse, error) {
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
	count := dao_redis.CartProductCount(context.Background(), cartId, in.UserId)

	return &product.CartProductCountResponse{Total: count}, nil
}
