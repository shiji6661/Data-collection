package logic

import (
	"context"
	"errors"
	"order_srv/dao/dao_mysql"
	"order_srv/dao/dao_redis"
	"order_srv/pkg"
	"order_srv/proto_order/order"

	"strconv"
)

func CreateOrderCart(in *order.CreateOrderCartRequest) (*order.CreateOrderCartResponse, error) {
	// 获取用户对应的购物车 ID
	cartId, err := dao_redis.GetCartIdByUserId(context.Background(), strconv.FormatInt(in.UserId, 10))
	if err != nil {
		return nil, errors.New("购物车ID获取失败")
	}

	// 检查指定商品是否存在于购物车
	if !dao_redis.IsCartProductExists(context.Background(), cartId, in.UserId, in.ProductId) {
		return nil, errors.New("商品不存在于购物车")
	}
	orderId, err := dao_mysql.OrderCreateByCart(in.UserId, in.ProductId)
	if err != nil {
		return nil, err
	}

	// 获取购物车中商品信息
	info := dao_redis.GetCartProductInfo(context.Background(), cartId, in.UserId, in.ProductId)
	//生成支付连接
	pay := pkg.NewAliPay()
	totalPrice := strconv.FormatInt(int64(orderId.TotalPrice), 10)
	url := pay.Pay(info.ProductName, orderId.OrderId, totalPrice)
	// 返回订单创建成功的响应
	return &order.CreateOrderCartResponse{
		OrderId: orderId.Id,
		Url:     url,
	}, nil
}
