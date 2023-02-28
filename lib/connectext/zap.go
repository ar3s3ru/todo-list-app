package connectext

import (
	"context"
	"errors"
	"time"

	"github.com/bufbuild/connect-go"
	"go.uber.org/zap"
)

func LoggingInterceptor(logger *zap.Logger) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			start := time.Now()

			resp, err := next(ctx, req)
			latency := time.Since(start)

			if err != nil && logger.Level().Enabled(zap.ErrorLevel) {
				fields := []zap.Field{
					zap.Duration("time", latency),
					zap.String("rpc", req.Spec().Procedure),
					zap.String("addr", req.Peer().Addr),
					zap.String("protocol", req.Peer().Protocol),
					zap.Error(err),
				}

				var connecterr *connect.Error
				if errors.As(err, &connecterr) {
					fields = append(fields, zap.String("code", connecterr.Code().String()))
				}

				logger.Error("request terminated", fields...)
			}

			if err == nil && logger.Level().Enabled(zap.InfoLevel) {
				logger.Info("request terminated",
					zap.Duration("time", latency),
					zap.String("rpc", req.Spec().Procedure),
					zap.String("addr", req.Peer().Addr),
					zap.String("protocol", req.Peer().Protocol),
				)
			}

			return resp, err
		})
	}
}
