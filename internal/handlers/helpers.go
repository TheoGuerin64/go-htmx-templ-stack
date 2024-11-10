package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func COMPONENT(ec echo.Context, code int, component templ.Component) error {
	ec.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	ec.Response().Status = code
	return component.Render(ec.Request().Context(), ec.Response().Writer)
}
