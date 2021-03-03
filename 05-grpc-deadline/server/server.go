package main

import (
	"context"
	"log"
	"net"
	"runtime"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "grpc-demo/05-grpc-deadline/proto"
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
	data := make(chan *pb.SimpleResponse, 1)
	go handle(ctx, req, data)
	select {
	case resp := <-data:
		return resp, nil
	case <-ctx.Done():
		return nil, status.Errorf(codes.Canceled, "client cancelled, abandoning")
	}
}

func handle(ctx context.Context, req *pb.SimpleRequest, data chan<- *pb.SimpleResponse) {
	select {
	case <-ctx.Done():
		log.Printf("ctx.Done: %v", ctx.Err())
		runtime.Goexit() // 超时后退出该 Go 携程
	case <-time.After(4 * time.Second): // 模拟耗时操作
		resp := pb.SimpleResponse{
			Code:  200,
			Value: "hello, " + req.Data,
		}
		// // 修改数据库前进行超时判断
		// if ctx.Err() == context.Canceled {
		// 	// 如果已经超时，则退出
		// }
		data <- &resp
	}
}
