package main

import (
	"big-red/services/big-red-service/proto"
	big_red_service "big-red/services/big-red-service/proto/big-red-service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "8090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	big_red_service.RegisterBigRedServiceServer(grpcServer, proto.Server{})
	grpcServer.Serve(lis)
}
