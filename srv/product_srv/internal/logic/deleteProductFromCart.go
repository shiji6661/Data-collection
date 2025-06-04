package logic

import (

	"context"
	"errors"
	"models/model_product/model_mysql"
	"product_srv/dao/dao_redis"
	"product_srv/proto_product/product"
	"strconv"
	"way/user"
)

// todo: 移除购物车中商品
func DeleteProductFromCart(in *product.RemoveFromCartRequest) (*product.RemoveFromCartResponse, error) {
	// 判断用户是否存在
	info, err := user.CommonGetUserIdInfo(in.UserId)
	if err != nil {
		return nil, err
	}
	if info.ID == 0 {
		return nil, errors.New("用户不存在！请重新登录！")
	}

	// 判断商品是否存在
	p, err := model_mysql.FindProductById(in.ProductId)
	if err != nil {
		return nil, errors.New("当前网络不稳定，请稍后再试！")
	}
	if p.ID == 0 {
		return nil, errors.New("商品不存在！请重新选择！")
	}

	cartId, err := dao_redis.GetCartIdByUserId(context.Background(), strconv.Itoa(int(in.UserId)))
	if err != nil {
		return nil, err
	}

	// 判断购物车中是否存在该商品
	exists := dao_redis.IsCartProductExists(context.Background(), cartId, in.UserId, in.ProductId)
	if exists == false {
		return nil, errors.New("购物车中不存在该商品！")
	}

	// 从购物车中移除此商品
	err = dao_redis.DeleteProductFromCart(context.Background(), cartId, in.UserId, in.ProductId)
	if err != nil {
		return nil, errors.New("移除失败！")
	}

	return &product.RemoveFromCartResponse{Success: "移除商品成功！"}, nil
}
