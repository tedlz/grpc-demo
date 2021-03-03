package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"

	pb "grpc-demo/02-grpc-server-stream/proto"
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
	// 默认单次接收消息最大长度为 1024*1024*4 bytes(4M)，单次发送消息最大长度为 math.MaxInt32 bytes
	// grpcServer := grpc.NewServer(grpc.MaxRecvMsgSize(1024*1024*4), grpc.MaxSendMsgSize(math.MaxInt32))
	grpcServer := grpc.NewServer()
	// 在 grpc 服务器注册我们的服务
	pb.RegisterStreamServerServer(grpcServer, &StreamService{})

	// 用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}

// StreamService 定义我们的服务
type StreamService struct{}

// Route 实现 Route 方法
func (s *StreamService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	resp := pb.SimpleResponse{
		Code:  200,
		Value: "hello, " + req.Data,
	}
	return &resp, nil
}

// ListValue 实现 ListValue 方法
func (s *StreamService) ListValue(req *pb.SimpleRequest, srv pb.StreamServer_ListValueServer) error {
	for i := 0; i < 15; i++ {
		// 向流中发送消息，默认每次发送消息的最大长度为 math.MaxInt32 bytes
		err := srv.Send(&pb.StreamResponse{StreamValue: req.Data + strconv.Itoa(i)})
		if err != nil {
			return err
		}
		log.Println(i)
		time.Sleep(1 * time.Second)
	}
	return nil
}
