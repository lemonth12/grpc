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
)

func DiscoverService(client *clientv3.Client) []string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	resp, err := client.Get(ctx, ServiceName, clientv3.WithPrefix())
	if err != nil {
		log.Fatalf("failed to discover service: %v", err)
	}
	addrArr := make([]string, 0)
	for _, kv := range resp.Kvs {
		addrArr = append(addrArr, string(kv.Value))
		log.Printf("Discovered service: %s -> %s", kv.Key, kv.Value)
	}

	return addrArr

	// Watch for changes  监控服务
	//rch := client.Watch(context.Background(), ServiceName, clientv3.WithPrefix())
	//for wresp := range rch {
	//	for _, ev := range wresp.Events {
	//		log.Printf("Watch event: %s %q : %q", ev.Type, ev.Kv.Key, ev.Kv.Value)
	//		//server[]
	//	}
	//}
}

func InvodiscoverService() []string {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{EtcdEndpoints},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("failed to connect to etcd: %v", err)
	}
	defer client.Close()

	addrArr := DiscoverService(client)
	return addrArr
}
