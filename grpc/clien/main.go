package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pd "grpctest/clien/grpc"
	"log"
	"time"
)

func main() {
	serverAddr := "localhost:50051"

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(), // 阻止 Dial 直到连接成功或失败
		grpc.WithTimeout(5 * time.Second),
	}
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, serverAddr, opts...)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	client := pd.NewHelloWorldClient(conn)
	hello, err := client.SayHello(context.Background(), &pd.HelloRequest{Name: "lemon"})
	if err != nil {
		log.Fatalf("failed to get task list: %v", err)
	}
	log.Printf("Task list: %v\n", hello)
}
