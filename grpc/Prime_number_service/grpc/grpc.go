package grpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"service/service"
)

type StatisServerr struct {
	UnimplementedStatisServer
}

func (s *StatisServerr) FindPrimeNum(ctx context.Context, in *ArrRequest) (*ArrResponse, error) {
	var ar *ArrResponse
	prime, err := service.GetPrime(in.Message)
	if err != nil {
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
