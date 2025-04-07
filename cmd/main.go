package main

func main() {
	//ctx := context.Background()
	//ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	//defer stop()
	//
	//ctx, err = logger.New(ctx)
	//if err != nil {
	//	logger.GetLoggerFromContext(ctx).Fatal(ctx, "can't initialize logger", zap.Error(err))
	//}
	//defer logger.GetLoggerFromContext(ctx).Sync()
	//
	//config, err := mainConfig.NewMainConfig()
	//if err != nil {
	//	logger.GetLoggerFromContext(ctx).Fatal(ctx, "can't read config", zap.Error(err))
	//}
	//
	//logger.GetLoggerFromContext(ctx).Info(ctx, "main-config: ", zap.Any("config", config))
	//
	//pgConn, err := postgres.NewPostgres(config.POSTGRES)
	//if err != nil {
	//	logger.GetLoggerFromContext(ctx).Fatal(ctx, "can't connect to db", zap.Error(err))
	//}
	//logger.GetLoggerFromContext(ctx).Info(ctx, "postgres", zap.Any("postgres", pgConn))
	//
	//redisConn := redis.NewRedisConn(config.REDIS)
	//logger.GetLoggerFromContext(ctx).Info(ctx, "redis", zap.Any("redis", redisConn))
	//
	//lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.UserServicePort))
	//if err != nil {
	//	logger.GetLoggerFromContext(ctx).Fatal(ctx, "failed to listen", zap.Error(err))
	//}
	//
	//server := grpc.NewServer()
	//svc := service2.NewUserService(pgConn)
	//userservice.RegisterUserServer(server, svc)
	//
	//logger.GetLoggerFromContext(ctx).Info(ctx, "server started", zap.String("port", strconv.Itoa(config.UserServicePort)))
	//go func() {
	//	if err := server.Serve(lis); err != nil {
	//		logger.GetLoggerFromContext(ctx).Fatal(ctx, "failed to serve", zap.Error(err))
	//	}
	//}()
	//
	//<-ctx.Done()
	//server.GracefulStop()
	//logger.GetLoggerFromContext(ctx).Info(ctx, "server stopped")
}
