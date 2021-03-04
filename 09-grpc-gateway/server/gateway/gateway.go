package gateway

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	pb "grpc-demo/09-grpc-gateway/proto"
	"grpc-demo/09-grpc-gateway/server/swagger"
)

// ProvideHTTP 把 grpc 服务转成 http 服务，让 grpc 同时支持 http
func ProvideHTTP(endpoint string, grpcServer *grpc.Server) *http.Server {
	ctx := context.Background()
	// 获取证书
	creds, err := credentials.NewClientTLSFromFile("./pkg/tls/server.pem", "virtual.machine.com")
	if err != nil {
		log.Fatalf("failed to create TLS credentials %v\n", err)
	}
	// 添加证书
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	// 新建 gwmux，它是 grpc-gateway 的请求复用器，它将 http 请求与模式匹配，并调用相应的处理程序
	gwmux := runtime.NewServeMux()
	// 将服务的 http 处理程序注册到 gwmux，处理程序通过 endpoint 转发请求到 grpc 端点
	err = pb.RegisterSimpleHandlerFromEndpoint(ctx, gwmux, endpoint, dopts)
	if err != nil {
		log.Fatalf("register endpoint err: %v\n", err)
	}
	// 新建 mux，它是 http 的请求复用器
	mux := http.NewServeMux()
	// 注册 gwmux
	mux.Handle("/", gwmux)
	// 注册 swagger
	mux.HandleFunc("/swagger/", swagger.ServeSwaggerFile)
	swagger.ServeSwaggerUI(mux)
	log.Println(endpoint + " HTTP.Listing with TLS and token...")
	return &http.Server{
		Addr:      endpoint,
		Handler:   grpcHandlerFunc(grpcServer, mux),
		TLSConfig: getTLSConfig(),
	}
}

// grpcHandlerFunc 根据不同的请求重定向到指定的 Handler 处理
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.HasPrefix(r.Header.Get("Content-Type"), "application/grpc") {
			fmt.Println("grpcServer.ServeHTTP...")
			grpcServer.ServeHTTP(w, r)
		} else {
			fmt.Println("otherHandler.ServeHTTP...")
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

// getTLSConfig 获取 TLS 配置
func getTLSConfig() *tls.Config {
	cert, _ := ioutil.ReadFile("./pkg/tls/server.pem")
	key, _ := ioutil.ReadFile("./pkg/tls/server.key")
	var demoKeyPair *tls.Certificate
	pair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		grpclog.Fatalf("TLS Keypair err: %v\n", err)
	}
	demoKeyPair = &pair
	return &tls.Config{
		Certificates: []tls.Certificate{*demoKeyPair},
		NextProtos:   []string{http2.NextProtoTLS},
	}
}
