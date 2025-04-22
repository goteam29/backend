package interceptors

import (
	"api-repository/pkg/utils"
	"context"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
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
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "you need to provide metadata :(")
		}

		values := md["authorisation"]
		if len(values) == 0 {
			return nil, status.Error(codes.Unauthenticated, "please authorise!")
		}

		token := strings.TrimPrefix(values[0], "Bearer")
		if token == values[0] {
			return nil, status.Error(codes.Unauthenticated, "invalid token format")
		}

		_, err := jwt.Verify(token)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "unauthorised")
		}
		return handler(ctx, req)

	}
}
