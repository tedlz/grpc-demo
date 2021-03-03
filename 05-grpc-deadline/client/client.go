package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "grpc-demo/05-grpc-deadline/proto"
)

const (
	// Address 连接地址
	Address string = "127.0.0.1:8000"
)

var grpcClient pb.SimpleClient

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	// 建立 grpc 连接
	grpcClient = pb.NewSimpleClient(conn)
	route(context.Background(), 2)
}

// route 调用服务端 Route 方法
func route(ctx context.Context, deadlines time.Duration) {
	// 设置超时时间
	clientDeadline := time.Now().Add(deadlines * time.Second)
	ctx, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()

	// 创建发送结构体
	req := pb.SimpleRequest{Data: "grpc"}
	resp, err := grpcClient.Route(ctx, &req)
	if err != nil {
		// 获取错误状态
		state, ok := status.FromError(err)
		if ok {
			// 判断是否为调用超时
			if state.Code() == codes.DeadlineExceeded {
				log.Fatalln("Route timeout!")
			}
		}
		log.Fatalf("Call Route err: %v", err)
	}
	log.Println(resp.Value)
}
