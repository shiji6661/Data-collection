package logic

import (
	"Data-collection/srv/kuteng-RabbitMQ/SimlpePublish"
	"Data-collection/srv/order_srv/pkg"
	"Data-collection/srv/order_srv/proto_order/order"
	"Data-collection/way/product"
	"Data-collection/way/user"
	"errors"

	"strconv"
)

func CreateOrder(in *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	//查找用户是否存在
	users, err := user.CommonGetUserIdInfo(in.UserId)
	if err != nil {
		return nil, errors.New("用户信息查询失败")
	}
	if users == nil {
		return nil, errors.New("用户信息为空")
	}
	//查找商品是否存在
	products, err := product.GetCommonFindProductById(in.ProductId)
	if err != nil {
		return nil, errors.New("商品信息查询失败")
	}
	//判断商品库存是否充足
	stock := products.Stock - in.Num
	if stock < 0 {
		return nil, errors.New("商品库存不足")
	}

	//创建该商品订单
	OrderSn := uuid.New().String()
	finalPrice := float64(in.Num) * products.Price
	//进入rabbitmq队列
	orderMsg := map[string]interface{}{
		"OrderId":      OrderSn,
		"Uid":          in.UserId,
		"RealName":     users.UserName,
		"UserPhone":    users.UserPhone,
		"UserAddress":  users.UserAddress,
		"CartId":       products.CateId,
		"FreightPrice": float64(5 * in.Num), //运费5元一件
		"TotalNum":     in.Num,
		"TotalPrice":   finalPrice,
		"PayType":      in.Pay,
		"MerId":        products.MerId,
		"ProductId":    in.ProductId,
	}
	err = SimlpePublish.SimplePublishb(orderMsg)
	if err != nil {
		zap.L().Info("rabbitmq发送失败")
		return nil, errors.New("rabbitmq发送失败")
	}

	// 生成支付链接
	pay := pkg.NewAliPay()
	totalPrice := strconv.FormatFloat(finalPrice, 'f', -1, 64)
	url := pay.Pay(products.StoreName, OrderSn, totalPrice)
	return &order.CreateOrderResponse{
		OrderSn: OrderSn,
		Url:     url,
	}, nil
}
