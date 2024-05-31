package main

import (
	"encoding/json"
	"fmt"
	"server/etcd"
	gp "server/grpc"
)

type Data struct {
	num int
}

func main() {
	//服务发现,k8s可以解决
	addrArr := etcd.InvodiscoverService()
	//var connect *grpc.ClientConn
	connect := gp.Connect(addrArr[0])
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	marshal, err := json.Marshal(ints)
	if err != nil {
		fmt.Println("转换失败", err)
	}
	statisGrpc, err := gp.InvoStatisGrpc(connect, string(marshal))
	if err != nil {
		panic(err)
	}
	fmt.Println(statisGrpc)

}
