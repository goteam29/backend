package interceptors

import (
	"api-repository/pkg/utils"
	"context"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

func AuthInterceptor(jwt *utils.JWTManager) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {

		utils.GetSugaredLogger().Logf(zapcore.InfoLevel, "gprc unary server info: %v", info)
		if info.FullMethod == "/user.User/Register" || info.FullMethod == "/user.User/Login" {

		}

		return nil, nil

	}
}
