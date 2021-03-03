package main

import (
	"context"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"

	pb "grpc-demo/07-grpc-middleware/proto"
	"grpc-demo/07-grpc-middleware/server/middleware/auth"
	"grpc-demo/07-grpc-middleware/server/middleware/cred"
	"grpc-demo/07-grpc-middleware/server/middleware/recovery"
	"grpc-demo/07-grpc-middleware/server/middleware/zap"
)

const (
	Address string = "127.0.0.1:8000"
	Network string = "tcp"
)

func main() {
	// 监听本地端口
	lis, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	// 新建 grpc 服务器实例
	grpcServer := grpc.NewServer(
		cred.TLSInterceptor(),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_auth.StreamServerInterceptor(auth.Interceptor),
			grpc_zap.StreamServerInterceptor(zap.Interceptor()),
			grpc_recovery.StreamServerInterceptor(recovery.Interceptor()),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(auth.Interceptor),
			grpc_zap.UnaryServerInterceptor(zap.Interceptor()),
			grpc_recovery.UnaryServerInterceptor(recovery.Interceptor()),
		)),
	)

	// 在 grpc 服务器注册我们的服务
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})
	log.Println(Address + " net.Listing with TLS and token...")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}

type SimpleService struct{}

func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	resp := pb.SimpleResponse{
		Code:  200,
		Value: "hello, " + req.Data,
	}
	return &resp, nil
}
