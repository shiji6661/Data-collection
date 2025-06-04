package internal

import (
	"collection_srv/proto_collection/collection"
	"collection_srv/server"
	"google.golang.org/grpc"

)

func Register(ser *grpc.Server) {
	collection.RegisterCollectionServer(ser, server.ServerCollection{})
}
