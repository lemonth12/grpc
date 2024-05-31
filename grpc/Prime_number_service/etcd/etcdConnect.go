package etcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

const (
	EtcdEndpoints = "localhost:2379"
	ServiceName   = "my-service"
	ServiceAddr   = "localhost:50051"
)

func registerService(client *clientv3.Client) {
	//key := fmt.Sprintf("/services/%s/%s", serviceName, serviceAddr)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := client.Put(ctx, ServiceName, ServiceAddr)
	if err != nil {
		log.Fatalf("failed to register service: %v", err)
	}
	//log.Printf("Service registered: %s", key)
}

func InvoregisterService() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{EtcdEndpoints},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("failed to connect to etcd: %v", err)
	}
	defer client.Close()
	registerService(client)
}
