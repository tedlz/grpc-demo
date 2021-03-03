package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "grpc-demo/06-grpc-security/proto"
)

const (
	Address   string = "127.0.0.1:8000"
	Network   string = "tcp"
	AppID     string = "grpc_token"
	AppSecret string = "123456"
)

func main() {
	// 监听本地端口
	lis, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	// 从输入证书文件和密钥文件为服务端构造 TLS 凭证
	creds, err := credentials.NewServerTLSFromFile("./pkg/tls/server.pem", "./pkg/tls/server.key")
	if err != nil {
		log.Fatalf("failed to generate credentials: %v", err)
	}

	// 普通方法：一元拦截器
	// grpc.UnaryServerInterceptor 为一元拦截器，只会拦截简单 RPC 方法。流式 RPC 方法需要使用流式拦截器 grpc.StreamInterceptor 进行拦截
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// 拦截普通方法请求，验证 token
		err = Check(ctx)
		if err != nil {
			return
		}
		// 继续处理请求
		return handler(ctx, req)
	}

	// 新建 grpc 服务器实例，并开启 TLS 和 token 认证
	// grpc.Creds：返回一个ServerOption，用于设置服务器连接的凭证
	grpcServer := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(interceptor))
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
	// 添加拦截器后，方法里省略 token 认证
	// // 检测 token 是否有效
	// if err := Check(ctx); err != nil {
	// 	return nil, err
	// }
	resp := pb.SimpleResponse{
		Code:  200,
		Value: "hello, " + req.Data,
	}
	return &resp, nil
}

// Check 验证 token
func Check(ctx context.Context) error {
	// metadata.FromIncomingContext：从上下文中获取元数据
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "获取 token 失败")
	}
	var appID, appSecret string
	if val, ok := md["app_id"]; ok {
		appID = val[0]
	}
	if val, ok := md["app_secret"]; ok {
		appSecret = val[0]
	}
	if appID != AppID || appSecret != AppSecret {
		return status.Errorf(codes.Unauthenticated, "token 无效：app_id=%s, app_secret=%s", appID, appSecret)
	}
	return nil
}
