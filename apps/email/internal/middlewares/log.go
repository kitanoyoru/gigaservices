package middlewares

import (
	"context"

	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"
)

func LogMiddleware(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		logger.Infof("[request] %v", req.Endpoint())
		err := fn(ctx, req, resp)
		return err
	}
}
