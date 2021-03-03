package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"grpc-demo/06-grpc-security/pkg/auth"
	pb "grpc-demo/06-grpc-security/proto"
)

const (
	Address   string = "127.0.0.1:8000"
	AppID     string = "grpc_token"
	AppSecret string = "123456"
)

var grpcClient pb.SimpleClient

func main() {
	// 从输入的证书文件中为客户端构造 TLS 凭证
	creds, err := credentials.NewClientTLSFromFile("./pkg/tls/server.pem", "CommonName")
	if err != nil {
		log.Fatalf("failed to create TLS credentials: %v", err)
	}
	// 构建 token
	token := auth.Token{
		AppID:     AppID,
		AppSecret: AppSecret,
	}
	// 连接服务器
	// grpc.WithTransportCredentials：配置连接级别的安全凭证（例如，TLS/SSL），返回一个 DialOption，用于连接服务器。
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&token))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	// 建立 grpc 连接
	grpcClient = pb.NewSimpleClient(conn)
	route()
}

// route 调用服务端 Route 方法
func route() {
	// 创建发送的结构体
	req := pb.SimpleRequest{Data: "grpc"}

	resp, err := grpcClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("grpcClient.Route err: %v", err)
	}
	log.Println(resp)
}
