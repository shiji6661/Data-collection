package server

import (
	"context"
	"order_srv/internal/logic"
	"order_srv/proto_order/order"
)

type ServerOrder struct {
	order.UnimplementedOrderServer
}

func (s ServerOrder) CreateOrder(ctx context.Context, in *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	res, err := logic.CreateOrder(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s ServerOrder) CreateOrderCart(ctx context.Context, in *order.CreateOrderCartRequest) (*order.CreateOrderCartResponse, error) {
	res, err := logic.CreateOrderCart(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s ServerOrder) OrderPay(ctx context.Context, in *order.OrderPayRequest) (*order.OrderPayResponse, error) {
	res, err := logic.OrderPay(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s ServerOrder) CreateOrderBargain(ctx context.Context, in *order.CreateOrderBargainRequest) (*order.CreateOrderBargainResponse, error) {
	bargain, err := logic.CreateOrderBargain(in)
	if err != nil {
		return nil, err
	}
	return bargain, err
}

func (s ServerOrder) OrderPayBakeCall(ctx context.Context, in *order.OrderPayBakeCallRequest) (*order.OrderPayBakeCallResponse, error) {
	res, err := logic.OrderPayBakeCall(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s ServerOrder) CreateOrderMessage(ctx context.Context, in *order.CreateOrderMessageRequest) (*order.CreateOrderMessageResponse, error) {
	create, err := logic.RabbitMqToCreate(in)
	if err != nil {
		return nil, err
	}
	return create, nil
}
func (s ServerOrder) CreateOrCode(ctx context.Context, in *order.CreateOrCodeRequest) (*order.CreateOrCodeResponse, error) {
	res, err := logic.CreateCode(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s ServerOrder) MerchantCountAmount(ctx context.Context, in *order.MerchantCountAmountRequest) (*order.MerchantCountAmountResponse, error) {
	res, err := logic.MerchantCountAmount(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s ServerOrder) OrderCreate(ctx context.Context, in *order.OrderCreateRequest) (*order.OrderCreateResponse, error) {
	create, err := logic.OrderCreate(in)
	if err != nil {
		return nil, err
	}
	return create, err

}
