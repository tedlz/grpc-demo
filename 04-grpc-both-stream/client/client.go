package main

import (
	"context"
	"io"
	"log"
	"strconv"

	"google.golang.org/grpc"

	pb "grpc-demo/04-grpc-both-stream/proto"
)

const (
	Address string = "127.0.0.1:8000"
)

var streamClient pb.StreamClient

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	// 建立 grpc 连接
	streamClient = pb.NewStreamClient(conn)
	conversations()
}

// conversations 调用服务端的 Conversations 方法
func conversations() {
	// 调用服务端的 Conversations 方法，获取流
	stream, err := streamClient.Conversations(context.Background())
	if err != nil {
		log.Fatalf("streamClient.Conversations err: %v", err)
	}
	for i := 0; i < 5; i++ {
		err := stream.Send(&pb.StreamRequest{Question: "stream client rpc: " + strconv.Itoa(i)})
		if err != nil {
			log.Fatalf("stream.Send err: %v", err)
		}
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream.Recv err: %v", err)
		}
		// 打印返回值
		log.Println(resp.Answer)
	}
	// 最后关闭流
	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("stream.CloseSend err: %v", err)
	}
}
