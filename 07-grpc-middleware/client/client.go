package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"grpc-demo/07-grpc-middleware/client/auth"
	pb "grpc-demo/07-grpc-middleware/proto"
)

const (
	Address string = "127.0.0.1:8000"
)

var grpcClient pb.SimpleClient

func main() {
	// 从输入的证书文件中为客户端构造 TLS 凭证
	creds, err := credentials.NewClientTLSFromFile("./pkg/tls/server.pem", "CommonName")
	if err != nil {
		log.Fatalf("failed to create TLS credentials: %v", err)
	}

	// 构建 token
	token := auth.Token{Value: "bearer grpc.auth.token"}
	// 连接服务器
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
	// 创建发送结构体
	req := pb.SimpleRequest{Data: "grpc"}

	resp, err := grpcClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("grpcClient.Route err: %v", err)
	}
	log.Println(resp)
}
