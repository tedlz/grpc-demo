package zap

import (
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Interceptor 返回 zap.Logger 实例（把日志写到文件中）
func Interceptor() *zap.Logger {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:  "./server/log/debug.log",
		MaxSize:   1024, // MB
		LocalTime: true,
	})

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(config), w, zap.NewAtomicLevel())

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	grpc_zap.ReplaceGrpcLoggerV2(logger)
	return logger
}

// Interceptor 返回 zap.Logger 实例（把日志输出到控制台）
// func Interceptor() *zap.Logger {
// 	logger, err := zap.NewDevelopment()
// 	if err != nil {
// 		log.Fatalf("zap.NewDevelopment err: %v", err)
// 	}
// 	grpc_zap.ReplaceGrpcLoggerV2(logger)
// 	return logger
// }
