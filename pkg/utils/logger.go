package utils

import (
	"context"
	"log"

	"go.uber.org/zap"
)

const (
	Key       = "logger"
	RequestID = "requestID"
)

var sugaredLogger *zap.SugaredLogger

type Logger struct {
	l *zap.Logger
}

func New(ctx context.Context) (context.Context, error) {

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, Key, &Logger{logger})
	return ctx, nil
}

func NewSugaredLogger() {
	lg, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't create logger | err: %v", err)
	}
	sugaredLogger = lg.Sugar()
}

func GetSugaredLogger() *zap.SugaredLogger {
	return sugaredLogger
}

func GetLoggerFromContext(ctx context.Context) *Logger {
	return ctx.Value(Key).(*Logger)
}

func (l *Logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequestID) != nil {
		fields = append(fields, zap.String("requestID", ctx.Value(RequestID).(string)))
	}
	l.l.Info(msg, fields...)
}

func (l *Logger) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequestID) != nil {
		fields = append(fields, zap.String("requestID", ctx.Value(RequestID).(string)))
	}
	l.l.Fatal(msg, fields...)
}

func (l *Logger) Sync() error {
	return l.l.Sync()
}

// func Interceptor(ctx context.Context,
// 	req any,
// 	info *grpc.UnaryServerInfo,
// 	next grpc.UnaryHandler,
// ) (any, error) {
// 	ctx, err := New(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("Interceptor: failed to create logger: %w", err)
// 	}

// 	guid := uuid.New().String()
// 	ctx = context.WithValue(ctx, RequestID, guid)

// 	GetLoggerFromContext(ctx).Info(ctx, "request", zap.String("method", info.FullMethod))

// 	resp, err := next(ctx, req)
// 	if err != nil {
// 		GetLoggerFromContext(ctx).Info(ctx, "response_error",
// 			zap.String("method: ", info.FullMethod),
// 			zap.String("error: ", err.Error()))
// 	} else {
// 		GetLoggerFromContext(ctx).Info(ctx, "response",
// 			zap.Any("response: ", resp))
// 	}

// 	return resp, err
// }
