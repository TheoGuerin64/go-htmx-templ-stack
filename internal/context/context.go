package context

import (
	"project/internal/config"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type Context struct {
	echo.Context

	Settings *config.Settings
	DB       *bun.DB
}

func New(c echo.Context, settings *config.Settings, db *bun.DB) *Context {
	return &Context{
		Context:  c,
		Settings: settings,
		DB:       db,
	}
}

func Convert(ec echo.Context) *Context {
	if c, ok := ec.(*Context); ok {
		return c
	}
	panic("unable to convert echo.Context to *context.Context")
}
