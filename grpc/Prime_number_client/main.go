package main

import (
	"fmt"
	"server/etcd"
	gp "server/grpc"
)

func main() {
	//服务发现
	addrArr := etcd.InvodiscoverService()
	//var connect *grpc.ClientConn
	connect := gp.Connect(addrArr[0])
	statisGrpc, err := gp.InvoStatisGrpc(connect, "")
	if err != nil {
		panic(err)
	}
	fmt.Println(statisGrpc)

}
