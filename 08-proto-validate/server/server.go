package main

import (
	"context"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"

	pb "grpc-demo/08-proto-validate/proto"
	"grpc-demo/08-proto-validate/server/middleware/auth"
	"grpc-demo/08-proto-validate/server/middleware/cred"
	"grpc-demo/08-proto-validate/server/middleware/recovery"
	"grpc-demo/08-proto-validate/server/middleware/zap"
)

const (
	Address string = "0.0.0.0:8000"
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
			grpc_validator.StreamServerInterceptor(),
			grpc_auth.StreamServerInterceptor(auth.Interceptor),
			grpc_zap.StreamServerInterceptor(zap.Interceptor()),
			grpc_recovery.StreamServerInterceptor(recovery.Interceptor()),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			grpc_auth.UnaryServerInterceptor(auth.Interceptor),
			grpc_zap.UnaryServerInterceptor(zap.Interceptor()),
			grpc_recovery.UnaryServerInterceptor(recovery.Interceptor()),
		)),
	)

	// 在 grpc 服务器注册我们的服务
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})
	log.Println(Address + " net.Listing with TLS and token...")

	// 用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}

type SimpleService struct{}

func (s *SimpleService) Route(ctx context.Context, req *pb.InnerMessage) (*pb.OuterMessage, error) {
	resp := pb.OuterMessage{
		ImportantString: "hello grpc validator",
		Inner:           req,
	}
	return &resp, nil
}
