package middleware

import (
	"project/internal/config"
	"project/internal/context"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

func GetContextMiddleware(settings *config.Settings, db *bun.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			c := context.New(ec, settings, db)
			return next(c)
		}
	}
}
