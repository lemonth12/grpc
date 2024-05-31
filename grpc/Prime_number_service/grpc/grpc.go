package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"service/service"
)

type StatisServerr struct {
	UnimplementedStatisServer
}

func (s *StatisServerr) FindPrimeNum(ctx context.Context, in *ArrRequest) (*ArrResponse, error) {
	ar := &ArrResponse{}
	fmt.Println("来了")
	prime, err := service.GetPrime(in.Message)
	if err != nil {
		fmt.Println("调用错误", err)
		return ar, err
	}
	ar.Message = prime
	return &ArrResponse{Message: prime}, nil
}

func ConnectGrpc() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterStatisServer(s, &StatisServerr{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
