package controllers

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewRouter returns new router
func NewRouter(h *Handler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == os.Getenv("API_KEY"), nil
	}))

	e.GET("/", h.indexHandler)

	return e
}

func (h *Handler) indexHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
