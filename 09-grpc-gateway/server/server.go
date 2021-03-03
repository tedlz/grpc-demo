package main

import (
	"context"
	"crypto/tls"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"

	pb "grpc-demo/09-grpc-gateway/proto"
	"grpc-demo/09-grpc-gateway/server/gateway"
	"grpc-demo/09-grpc-gateway/server/middleware/auth"
	"grpc-demo/09-grpc-gateway/server/middleware/cred"
	"grpc-demo/09-grpc-gateway/server/middleware/recovery"
	"grpc-demo/09-grpc-gateway/server/middleware/zap"
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

	// 使用 gateway 把 grpcServer 转成 httpServer
	httpServer := gateway.ProvideHTTP(Address, grpcServer)
	if err = httpServer.Serve(tls.NewListener(lis, httpServer.TLSConfig)); err != nil {
		log.Fatalf("ListenAndServe err: %v\n", err)
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
