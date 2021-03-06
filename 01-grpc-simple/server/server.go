package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpc-demo/01-grpc-simple/proto"
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
	// 在 grpc 服务器注册我们的服务
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})
	// 用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}

// SimpleService 定义我们的服务
type SimpleService struct{}

// Route 实现 Route 方法
func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	resp := pb.SimpleResponse{
		Code:  200,
		Value: "hello, " + req.Data,
	}
	return &resp, nil
}
