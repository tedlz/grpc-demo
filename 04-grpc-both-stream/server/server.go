package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	pb "grpc-demo/04-grpc-both-stream/proto"
)

const (
	// Address 监听地址
	Address string = "127.0.0.1:8000"
	// Network 网络通信协议
	Network string = "tcp"
)

func main() {
	// 监听本地端口
	lis, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")

	// 新建 grpc 服务器实例
	grpcServer := grpc.NewServer()
	// 向 grpc 服务器注册我们的服务
	pb.RegisterStreamServer(grpcServer, &StreamService{})

	// 用 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}

// StreamService 定义我们的服务
type StreamService struct{}

func (s *StreamService) Conversations(srv pb.Stream_ConversationsServer) error {
	i := 1
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		answer := fmt.Sprintf("from stream server answer: the %s question is %s", strconv.Itoa(i), req.Question)
		err = srv.Send(&pb.StreamResponse{Answer: answer})
		if err != nil {
			return err
		}
		i++
		log.Printf("from stream client question: %s", req.Question)
	}
}
