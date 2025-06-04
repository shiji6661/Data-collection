package internal

import (
	"google.golang.org/grpc"
	"weikang/Data-collection/srv/collection_srv/proto_collection/collection"
	"weikang/Data-collection/srv/collection_srv/server"
)

func Register(ser *grpc.Server) {
	collection.RegisterCollectionServer(ser, server.ServerCollection{})
}
