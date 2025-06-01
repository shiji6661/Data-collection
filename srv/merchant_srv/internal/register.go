package internal

import (
	"google.golang.org/grpc"
	"merchant_srv/proto_merchant/merchant"
	server2 "merchant_srv/server"
)

func RegisterMerchant(server grpc.ServiceRegistrar) {
	merchant.RegisterMerchantServer(server, server2.ServerMerchant{})
}
