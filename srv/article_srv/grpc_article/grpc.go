package grpc_article

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

// RegisterGrpc 注册grpc服务
func RegisterGrpc(call func(grpc *grpc.Server)) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s", "127.0.0.1:8005"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	call(s)
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
