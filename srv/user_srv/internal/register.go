package internal

import (
	"google.golang.org/grpc"
	"user_srv/proto_user/user"
	"user_srv/server"
)

func Register(ser *grpc.Server) {
	user.RegisterUserServer(ser, server.ServerUser{})
}
