package config

import (
	"context"
	"log/slog"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func GetLoggerConfig(logger *slog.Logger) echomiddleware.RequestLoggerConfig {
	return echomiddleware.RequestLoggerConfig{
		LogURI:      true,
		LogMethod:   true,
		LogStatus:   true,
		LogLatency:  true,
		LogRemoteIP: true,
		LogError:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.String("method", v.Method),
					slog.Int("status", v.Status),
					slog.String("latency", v.Latency.String()),
					slog.String("remote_ip", v.RemoteIP),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.String("method", v.Method),
					slog.Int("status", v.Status),
					slog.String("latency", v.Latency.String()),
					slog.String("remote_ip", v.RemoteIP),
					slog.String("error", v.Error.Error()),
				)
			}
			return nil
		},
	}
}

var SlogHandlerOptions = slog.HandlerOptions{
	Level: slog.LevelInfo,
}

var RecoverConfig = echomiddleware.RecoverConfig{
	DisableErrorHandler: true,
}

func SentryClientOptions(dsn string) sentry.ClientOptions {
	return sentry.ClientOptions{
		Dsn:              dsn,
		AttachStacktrace: true,
	}
}

var SentryEchoOptions sentryecho.Options = sentryecho.Options{
	Repanic: true,
}
