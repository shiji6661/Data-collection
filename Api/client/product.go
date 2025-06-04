package client

import (
	"context"
	"google.golang.org/grpc"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
)

type HandlerProduct func(ctx context.Context, in product.ProductClient) (interface{}, error)

func ProductClient(ctx context.Context, handler HandlerProduct) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8002", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := product.NewProductClient(dial)
	res, err := handler(ctx, client)
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	return res, nil
}
