package handler

import "github.com/labstack/echo/v4"

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Hello(c echo.Context) error {
	return c.String(200, "hello")
}
