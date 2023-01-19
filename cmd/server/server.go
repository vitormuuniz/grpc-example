package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-stream-live/pb"
	"grpc-stream-live/servers"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterMathServiceServer(grpcServer, &servers.Math{})
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
