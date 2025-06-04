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

// TODO: 修改购物车中商品的数量
func UpdateProductCart(in *product.UpdateCartRequest) (*product.UpdateCartResponse, error) {
	// 判断用户是否存在
	info, err := user.CommonGetUserIdInfo(in.UserId)
	if err != nil {
		return nil, err
	}
	if info.ID == 0 {
		return nil, errors.New("用户不存在！请重新登录！")
	}

	userId := strconv.Itoa(int(in.UserId))
	cartId, err := dao_redis.GetCartIdByUserId(context.Background(), userId)
	if err != nil {
		return nil, err
	}

	id, err := model_mysql.FindProductById(in.ProductId)
	if err != nil {
		return nil, errors.New("当前网络错误！请稍后重试！")
	}
	if id.ID == 0 {
		return nil, errors.New("商品不存在！请重新选择！")
	}

	// 判断购物车中是否存在该商品
	exists := dao_redis.IsCartProductExists(context.Background(), cartId, in.UserId, in.ProductId)
	if exists == false {
		return nil, errors.New("购物车中不存在该商品！")
	}

	// 判断要修改的数量是否大于库存
	num := dao_redis.GetProductsFromRedis(context.Background(), in.ProductId)
	if num < in.Num {
		return nil, errors.New("商品库存不足！")
	}
	// 修改购物车中商品的数量
	err = dao_redis.UpdateProductCart(context.Background(), cartId, in.UserId, in.ProductId, in.Num)
	if err != nil {
		return nil, errors.New("修改数量失败！")
	}
	return &product.UpdateCartResponse{Success: "修改数量成功！"}, nil
}
