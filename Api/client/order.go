package client

import (
	"context"
	"google.golang.org/grpc"
	"order_srv/proto_order/order"
)

type HandlerOrder func(ctx context.Context, in order.OrderClient) (interface{}, error)

func OrderClient(ctx context.Context, handler HandlerOrder) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8003", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := order.NewOrderClient(dial)
	res, err := handler(ctx, client)
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	return res, nil
}
