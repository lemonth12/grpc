package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pd "grpctest/service/grpc"
	"log"
	"net"
)

type HelloWorldServer struct {
	pd.UnimplementedHelloWorldServer
}

func (s *HelloWorldServer) SayHello(ctx context.Context, empty *pd.HelloRequest) (*pd.HelloResponse, error) {
	fmt.Println("diayong")
	return &pd.HelloResponse{Message: empty.Name}, nil
}

func main() {
	var arr []int
	arr = []int{1, 2, 3}
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pd.RegisterHelloWorldServer(s, &HelloWorldServer{})
	//pd.RegisterTaskListServiceServer(s, &respServer{})
	log.Println("Server running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println(arr)
}
