package interceptors

import (
	"context"
	"github.com/yanshicheng/opsx-common/handler/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
)

func ServerErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		logx.Info("ServerErrorInterceptor:", err)
		return resp, errorx.FromError(err).Err()
	}
}
