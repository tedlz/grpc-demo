package auth

import (
	"context"
	"errors"
	"log"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TokenInfo 用户信息
type TokenInfo struct {
	ID    string
	Roles []string
}

// Interceptor 认证拦截器，对以 authorization 为头部，形式为 bearer token 的 token 进行验证
func Interceptor(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	tokenInfo, err := parseToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, " %v", err)
	}
	// 使用 context.WithValue 添加了值后，可以用 Value(key) 方法获取值
	newCtx := context.WithValue(ctx, tokenInfo.ID, tokenInfo)
	log.Println(newCtx.Value(tokenInfo.ID))
	return newCtx, nil
}

func parseToken(token string) (TokenInfo, error) {
	var tokenInfo TokenInfo
	if token == "grpc.auth.token" {
		tokenInfo.ID = "1"
		tokenInfo.Roles = []string{"admin"}
		return tokenInfo, nil
	}
	return tokenInfo, errors.New("Token 无效：bearer " + token)
}

// 从 token 中获取用户唯一标识
func userClaimFromToken(tokenInfo TokenInfo) string {
	return tokenInfo.ID
}
