package logic

import (
	"Data-collection/way/user"
	"context"
	"errors"
	"weikang/Data-collection/srv/product_srv/dao/dao_redis"
	"weikang/Data-collection/srv/product_srv/proto_product/product"

	"strconv"
)

// TODO: 购物车商品列表展示
func ProductCartList(in *product.CartProductListRequest) (*product.CartProductListResponse, error) {
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
	list := dao_redis.ProductCartList(context.Background(), cartId, in.UserId)

	return &product.CartProductListResponse{List: list}, nil
}
