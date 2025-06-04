package logic

import (
	"Data-collection/way/user"
	"context"
	"errors"
	"weikang/Data-collection/srv/product_srv/dao/dao_redis"
	"weikang/Data-collection/srv/product_srv/proto_product/product"

	"strconv"
)

// TODO: 用户清空购物车
func ClearProductCart(in *product.ClearCartRequest) (*product.ClearCartResponse, error) {
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
	err = dao_redis.ClearProductCart(context.Background(), cartId, in.UserId)
	if err != nil {
		return nil, errors.New("清空购物车失败！")
	}

	return &product.ClearCartResponse{Success: "清空购物车成功！"}, nil
}
