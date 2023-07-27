package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	grpccontext "github.com/kitanoyoru/gigaservices/pkg/grpc/context"
	"github.com/sirupsen/logrus"
)

func NewRequestLogger() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		ctx = grpccontext.SetRequestUniqueID(ctx)
		reqId := grpccontext.GetRequestID(ctx)

		logrus.WithFields(logrus.Fields{
			"method":     info.FullMethod,
			"request_id": reqId,
		}).Info("GRPC request")

		res, err := handler(ctx, req)

		logrus.WithFields(logrus.Fields{
			"method":     info.FullMethod,
			"code":       status.Code(err),
			"request_id": reqId,
		}).Info("GRPC response")

		return res, err
	}
}
