package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "grpc-demo/01-grpc-simple/proto"
)

const (
	// Address 连接地址
	Address string = "127.0.0.1:8000"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	// 建立 grpc 连接
	grpcClient := pb.NewSimpleClient(conn)
	// 创建发送的结构体
	req := pb.SimpleRequest{Data: "grpc"}
	// 调用我们的服务（ Route方法 ）
	// 同时传入了一个 context.Context，在有需要时可以让我们改变 RPC 的行为，比如超时 / 取消一个正在运行的 RPC
	resp, err := grpcClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("grpcClient.Route err: %v", err)
	}
	// 打印返回值
	log.Println(resp)
}
