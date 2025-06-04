package internal

import (
	"google.golang.org/grpc"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
	"weikang/Data-collection/srv/user_srv/server"
)

func Register(ser *grpc.Server) {
	user.RegisterUserServer(ser, server.ServerUser{})
}
