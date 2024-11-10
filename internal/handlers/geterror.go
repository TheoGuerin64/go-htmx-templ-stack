package handlers

import (
	"net/http"
	"project/internal/templates"

	"github.com/labstack/echo/v4"
)

type ErrorHandLer struct{}

func NewErrorHandler() *ErrorHandLer {
	return &ErrorHandLer{}
}

func (h *ErrorHandLer) ServeHTTPError(err error, ec echo.Context) {
	var code int
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	} else {
		code = http.StatusInternalServerError
	}

	component := templates.Layout(templates.Error(code), "Error")
	_ = COMPONENT(ec, code, component)
}
