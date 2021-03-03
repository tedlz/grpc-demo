package main

import (
	"context"
	"io"
	"log"
	"strconv"

	"google.golang.org/grpc"

	pb "grpc-demo/03-grpc-client-stream/proto"
)

const (
	// Address 连接地址
	Address string = "127.0.0.1:8000"
)

var streamClient pb.StreamClientClient

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	// 建立 grpc 连接
	streamClient = pb.NewStreamClientClient(conn)
	route()
	routeList()
}

// route 调用服务端 Route 方法
func route() {
	req := pb.SimpleRequest{Data: "grpc"}

	resp, err := streamClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("streamClient.Route err: %v", err)
	}
	log.Println(resp)
}

// routeList 调用服务端 RouteList 方法
func routeList() {
	// 调用服务端 RouteList 方法，获取流
	stream, err := streamClient.RouteList(context.Background())
	if err != nil {
		log.Fatalf("streamClient.RouteList err: %v", err)
	}
	for i := 0; i < 5; i++ {
		// 向流中发送消息
		err := stream.Send(&pb.StreamRequest{StreamData: "stream client rpc " + strconv.Itoa(i)})
		// 发送也要检测 EOF，当服务端在消息没接收完前主动调用 SendAndClose() 关闭 stream，此时客户端还执行 Send()，则会返回 EOF 错误
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream.Send err: %v", err)
		}
	}
	// 关闭流并获取返回的消息
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("stream.CloseAndRecv err: %v", err)
	}
	log.Println(resp)
}
