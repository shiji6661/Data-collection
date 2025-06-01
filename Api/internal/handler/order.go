package handler

import (
	"Api/client"
	"context"
	"order_srv/proto_order/order"
)

// todo 商品下单
func CreateOrder(ctx context.Context, i *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	orderClient, err := client.OrderClient(ctx, func(ctx context.Context, in order.OrderClient) (interface{}, error) {
		createOrder, err := in.CreateOrder(ctx, i)
		if err != nil {
			return nil, err
		}
		return createOrder, nil
	})
	if err != nil {
		return nil, err
	}
	return orderClient.(*order.CreateOrderResponse), nil
}

// todo 商品下单(购物车)
func CreateOrderCart(ctx context.Context, i *order.CreateOrderCartRequest) (*order.CreateOrderCartResponse, error) {
	orderClient, err := client.OrderClient(ctx, func(ctx context.Context, in order.OrderClient) (interface{}, error) {
		createOrder, err := in.CreateOrderCart(ctx, i)
		if err != nil {
			return nil, err
		}
		return createOrder, nil
	})
	if err != nil {
		return nil, err
	}
	return orderClient.(*order.CreateOrderCartResponse), nil
}

// todo 订单支付(异步)
func OrderPay(ctx context.Context, i *order.OrderPayRequest) (*order.OrderPayResponse, error) {
	orderClient, err := client.OrderClient(ctx, func(ctx context.Context, in order.OrderClient) (interface{}, error) {
		createOrder, err := in.OrderPay(ctx, i)
		if err != nil {
			return nil, err
		}
		return createOrder, nil
	})
	if err != nil {
		return nil, err
	}
	return orderClient.(*order.OrderPayResponse), nil
}

func CreateOrderBargain(ctx context.Context, i *order.CreateOrderBargainRequest) (*order.CreateOrderBargainResponse, error) {
	orderClient, err := client.OrderClient(ctx, func(ctx context.Context, in order.OrderClient) (interface{}, error) {
		bargain, err := in.CreateOrderBargain(ctx, i)
		if err != nil {
			return nil, err
		}
		return bargain, err
	})
	if err != nil {
		return nil, err
	}
	return orderClient.(*order.CreateOrderBargainResponse), nil
}

// todo 支付回调
func OrderCallBack(ctx context.Context, i *order.OrderPayBakeCallRequest) (*order.OrderPayBakeCallResponse, error) {
	orderClient, err := client.OrderClient(ctx, func(ctx context.Context, in order.OrderClient) (interface{}, error) {
		createOrder, err := in.OrderPayBakeCall(ctx, i)
		if err != nil {
			return nil, err
		}
		return createOrder, nil

	})
	if err != nil {
		return nil, err
	}
	return orderClient.(*order.OrderPayBakeCallResponse), nil

}

// todo 消费队列进行下单

func RabbitMqToCreate(ctx context.Context, i *order.CreateOrderMessageRequest) (*order.CreateOrderMessageResponse, error) {
	orderClient, err := client.OrderClient(ctx, func(ctx context.Context, in order.OrderClient) (interface{}, error) {
		create, err := in.CreateOrderMessage(ctx, i)
		if err != nil {
			return nil, err
		}
		return create, nil
	})
	if err != nil {
		return nil, err
	}
	return orderClient.(*order.CreateOrderMessageResponse), nil
}
func OrderCreateSpike(ctx context.Context, i *order.OrderCreateRequest) (*order.OrderCreateResponse, error) {
	orderClient, err := client.OrderClient(ctx, func(ctx context.Context, in order.OrderClient) (interface{}, error) {
		create, err := in.OrderCreate(ctx, i)
		if err != nil {
			return nil, err
		}
		return create, err

	})
	if err != nil {
		return nil, err
	}
	return orderClient.(*order.OrderCreateResponse), nil

}

// 生成二维码
func CreateOrderQrCode(ctx context.Context, i *order.CreateOrCodeRequest) (*order.CreateOrCodeResponse, error) {
	orderClient, err := client.OrderClient(ctx, func(ctx context.Context, in order.OrderClient) (interface{}, error) {
		createOrder, err := in.CreateOrCode(ctx, i)
		if err != nil {
			return nil, err
		}
		return createOrder, nil
	})
	if err != nil {
		return nil, err
	}

	return orderClient.(*order.CreateOrCodeResponse), nil
}

// 商家查看销售额
func MerchantSales(ctx context.Context, i *order.MerchantCountAmountRequest) (*order.MerchantCountAmountResponse, error) {
	orderClient, err := client.OrderClient(ctx, func(ctx context.Context, in order.OrderClient) (interface{}, error) {
		createOrder, err := in.MerchantCountAmount(ctx, i)
		if err != nil {
			return nil, err
		}
		return createOrder, nil
	})
	if err != nil {
		return nil, err
	}
	return orderClient.(*order.MerchantCountAmountResponse), nil

}
