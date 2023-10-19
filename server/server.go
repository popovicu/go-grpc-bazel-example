package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"popovicu/foo"

	grpc "google.golang.org/grpc"
)

type sampleServer struct{}

func (s sampleServer) DoFoo(ctx context.Context, req *foo.FooRequest) (*foo.FooResponse, error) {
	return &foo.FooResponse{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:9876"))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	foo.RegisterEchoServiceServer(grpcServer, sampleServer{})

	fmt.Println("Starting")
	grpcServer.Serve(lis)

	fmt.Println("Hello!")
}
