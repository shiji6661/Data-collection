package logic

import (
	order2 "way/order"
	"way/product"
	"way/user"

	"errors"
	"models/model_user/model_mysql"
	"order_srv/dao/dao_mysql"
	"order_srv/pkg"
	"order_srv/proto_order/order"
	"strconv"
)

func OrderPay(in *order.OrderPayRequest) (*order.OrderPayResponse, error) {
	//查找用户是否存在
	users, err := user.CommonGetUserIdInfo(in.UserId)
	if err != nil {
		return nil, errors.New("用户信息查询失败")
	}
	if users == nil {
		return nil, errors.New("用户信息为空")
	}
	//查找订单是否存在
	ord, err := order2.FindByOrderId(in.OrderId)
	if err != nil {
		return nil, errors.New("订单信息查找失败")
	}
	if ord == nil {
		return nil, errors.New("订单信息为空")
	}
	//查找商品信息
	products, err := product.GetCommonFindProductById(ord.ProductId)
	if err != nil {
		return nil, errors.New("商品信息查询失败")
	}
	//支付方式
	err = dao_mysql.UpdateProductPay(in.OrderId, in.Pay)
	if err != nil {
		return nil, errors.New("支付方式修改失败")
	}

	price := strconv.FormatFloat(products.Price, 'f', -1, 64)
	pay := pkg.NewAliPay().Pay(products.StoreName, ord.OrderId, price)

	pf, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return nil, errors.New("转换出错")
	}
	member := &model_mysql.Member{}
	err = member.AddPointsByConsumption(in.UserId, pf, ord.OrderId)

	return &order.OrderPayResponse{
		Url: pay,
	}, nil

}
