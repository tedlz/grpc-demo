package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"

	pb "grpc-demo/02-grpc-server-stream/proto"
)

const (
	// Address 连接地址
	Address string = "127.0.0.1:8000"
)

var grpcClient pb.StreamServerClient

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	// 建立 grpc 连接
	grpcClient = pb.NewStreamServerClient(conn)
	route()
	listValue()
}

// route 调用服务端的 Route 方法
func route() {
	req := pb.SimpleRequest{Data: "grpc"}

	resp, err := grpcClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("grpcClient.Route err: %v", err)
	}
	log.Println(resp)
}

// listValue 调用服务端的 ListValue 方法
func listValue() {
	// 创建发送的结构体
	req := pb.SimpleRequest{Data: "stream server grpc"}

	// 调用我们的服务（ ListValue 方法 ）
	stream, err := grpcClient.ListValue(context.Background(), &req)
	if err != nil {
		log.Fatalf("grpcClient.ListValue err: %v", err)
	}
	for {
		// Recv() 方法接收服务端的消息
		resp, err := stream.Recv()
		// 判断消息流是否已经结束
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream.Recv err: %v", err)
		}
		// 打印返回值
		log.Println(resp.StreamValue)
		break
	}
	// 可以使用 CloseSend() 关闭 stream，这样服务端就不会继续产生流消息
	// 调用 CloseSend() 后，若继续调用 Recv()，会重新激活 stream，接着之前的结果获取消息
	// 这能完美解决客户端暂停 -> 继续获取数据的操作
	stream.CloseSend()
}
