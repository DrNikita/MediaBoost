package context

import (
	"context"

	"go.uber.org/zap"
)

func CtxWithLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, "logger", logger)
}

func FromContext(ctx context.Context) *zap.SugaredLogger {
	if logger, ok := ctx.Value("loggerKey").(*zap.SugaredLogger); ok {
		return logger
	}

	// TODO: add default logger
	return nil
}
