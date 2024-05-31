package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var serviceConfig = `{
"loadBalancingPolicy":"round_robin"
}`

func Connect(serverAddr string) *grpc.ClientConn {

	fmt.Println(serverAddr, "**************")

	opts := []grpc.DialOption{
		//传输方式，这里没有进行加密
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//grpc.WithBlock(), // 阻止 Dial 直到连接成功或失败
		grpc.WithDefaultServiceConfig(serviceConfig),
	}
	//ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, "localhost:50051", opts...)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	//defer conn.Close()
	return conn
}

func InvoStatisGrpc(conn *grpc.ClientConn, arrJson string) (*ArrResponse, error) {
	client := NewStatisClient(conn)
	hello, err := client.FindPrimeNum(context.Background(), &ArrRequest{Message: arrJson})
	if err != nil {
		log.Fatalf("failed to get task list: %v", err)
		return hello, err
	}
	log.Printf("Task list: %v\n", hello)
	return hello, nil
}
