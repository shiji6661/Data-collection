package client

import (
	"collection_srv/proto_collection/collection"
	"context"
	"google.golang.org/grpc"

)

type HandlerCollection func(ctx context.Context, in collection.CollectionClient) (interface{}, error)

func CollectionClient(ctx context.Context, handler HandlerCollection) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8006", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := collection.NewCollectionClient(dial)
	res, err := handler(ctx, client)
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	return res, nil
}
