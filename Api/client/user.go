package client

import (
	"context"
	"google.golang.org/grpc"
	"user_srv/proto_user/user"
)

type HandlerUser func(ctx context.Context, in user.UserClient) (interface{}, error)

func UserClient(ctx context.Context, handler HandlerUser) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8001", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := user.NewUserClient(dial)
	res, err := handler(ctx, client)
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	return res, nil
}
