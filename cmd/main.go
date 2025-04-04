package main

import (
	mainConfig "api-repository/internal/config"
	userservice "api-repository/pkg/api/user-service"
	"api-repository/pkg/db/postgres"
	"api-repository/pkg/db/redis"
	logger "api-repository/pkg/utils"
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	userservice.UnimplementedUserServer
}

func New() *UserService {
	return &UserService{}
}

func (u *UserService) Get(ctx context.Context, req *userservice.Request) (*userservice.Reply, error) {
	if req.Message == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request cannot be empty")
	}

	return &userservice.Reply{Message: req.Message}, nil
}

const port = "50050"

func main() {
	ctx := context.Background()
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	ctx, err := logger.New(ctx)
	if err != nil {
		logger.GetLoggerFromContext(ctx).Fatal(ctx, "can't initialize logger", zap.Error(err))
	}
	defer logger.GetLoggerFromContext(ctx).Sync()

	config, err := mainConfig.NewMainConfig()
	if err != nil {
		logger.GetLoggerFromContext(ctx).Fatal(ctx, "can't read config", zap.Error(err))
	}
	logger.GetLoggerFromContext(ctx).Info(ctx, "config", zap.Any("config", config))

	pgConn, err := postgres.NewPostgres(config.PgConf)
	if err != nil {
		logger.GetLoggerFromContext(ctx).Fatal(ctx, "can't connect to db", zap.Error(err))
	}
	logger.GetLoggerFromContext(ctx).Info(ctx, "postgres", zap.Any("postgres", pgConn))

	redisConn := redis.NewRedisConn(config.RedisConf)
	logger.GetLoggerFromContext(ctx).Info(ctx, "redis", zap.Any("redis", redisConn))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		logger.GetLoggerFromContext(ctx).Fatal(ctx, "failed to listen", zap.Error(err))
	}

	server := grpc.NewServer()
	service := New()
	userservice.RegisterUserServer(server, service)

	logger.GetLoggerFromContext(ctx).Info(ctx, "server started", zap.String("port", port))
	go func() {
		if err := server.Serve(lis); err != nil {
			logger.GetLoggerFromContext(ctx).Fatal(ctx, "failed to serve", zap.Error(err))
		}
	}()

	<-ctx.Done()
	server.GracefulStop()
	logger.GetLoggerFromContext(ctx).Info(ctx, "server stopped")
}
