package context

import (
	"context"

	"github.com/google/uuid"
)

func GetRequestID(ctx context.Context) string {
	val := ctx.Value(RequestIdContextKey)
	if val == nil {
		return ""
	}

	id, ok := val.(string)
	if !ok {
		return id
	}

	return id
}

func SetRequestUniqueID(ctx context.Context) context.Context {
	if id := GetRequestID(ctx); id != "" {
		return ctx
	}

	return context.WithValue(ctx, RequestIdContextKey, uuid.New().String())
}
