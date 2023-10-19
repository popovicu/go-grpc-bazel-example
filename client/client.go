package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"popovicu/foo"

	grpc "google.golang.org/grpc"
	insecure "google.golang.org/grpc/credentials/insecure"
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial("localhost:9876", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer conn.Close()

	client := foo.NewEchoServiceClient(conn)

	resp, err := client.DoFoo(context.Background(), &foo.FooRequest{})

	if err != nil {
	   log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("Response: %v\n", resp)
}
