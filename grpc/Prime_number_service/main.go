package main

import (
	"service/etcd"
	"service/grpc"
)

func main() {
	//服务注册
	etcd.InvoregisterService()
	//grpc连接
	grpc.ConnectGrpc()

}
