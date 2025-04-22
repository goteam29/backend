package interceptors

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"time"
)

func LoggingInterceptor(l *zap.SugaredLogger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		t := time.Now()
		l.Logf(0, "\n Request: %s | Time: %s \n FullInfo %v", info.FullMethod, t, info)

		resp, err := handler(ctx, req)

		if err != nil {
			l.Logf(zapcore.ErrorLevel, "Response: %s | Duration: %s | Err: %v", resp, time.Since(t), err)
		} else {
			l.Logf(zapcore.InfoLevel, "Response: %s | Duration: %s | Successfully completed ", resp, time.Since(t))
		}
		return resp, err
	}
}
