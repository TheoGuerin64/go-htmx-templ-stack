package handlers

import (
	"net/http"
	"project/internal/context"
	"project/internal/templates"

	"github.com/labstack/echo/v4"
)

type HomeHandLer struct{}

func NewHomeHandler() *HomeHandLer {
	return &HomeHandLer{}
}

func (h *HomeHandLer) ServeHTTP(ec echo.Context) error {
	c := context.Convert(ec)
	component := templates.Layout(templates.Home(), "Home")
	return COMPONENT(c, http.StatusOK, component)
}
