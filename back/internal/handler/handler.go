package handler

import (
	"app/internal/db"
	"app/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	conn *db.ConnectionDB
}

func NewHandler(connDB *db.ConnectionDB) *handler {
	return &handler{
		conn: connDB,
	}
}

func (h *handler) Hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}

func (h *handler) Test(c echo.Context) error {
	res, err := h.conn.GetTestName()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func (h *handler) CreateTest(c echo.Context) error {
	data := &model.TestData{}
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := h.conn.NewTestName(data.Name); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "OK")
}

func (h *handler) Close() error {
	return h.conn.Close()
}
