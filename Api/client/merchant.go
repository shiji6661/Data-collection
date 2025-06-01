package client

import (
	"context"
	"google.golang.org/grpc"
	"merchant_srv/proto_merchant/merchant"
)

type HandlerMerchant func(ctx context.Context, in merchant.MerchantClient) (interface{}, error)

func MerchantClient(ctx context.Context, handler HandlerMerchant) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8004", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := merchant.NewMerchantClient(dial)
	res, err := handler(ctx, client)
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	return res, nil
}
