package interceptors

import (
	"context"
	"errors"
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

		var cancel context.CancelFunc

		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()

		resp, err := handler(ctx, req)

		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				l.Logf(zapcore.ErrorLevel, "Timeout reached")
			}
			l.Logf(zapcore.ErrorLevel, "Response: %s | Duration: %s | Err: %v", resp, time.Since(t), err)
			return resp, err
		} else {
			l.Logf(zapcore.InfoLevel, "Response: %s | Duration: %s | Successfully completed ", resp, time.Since(t))
			return resp, err
		}
	}
}
