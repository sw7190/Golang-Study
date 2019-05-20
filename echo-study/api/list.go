package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func getList(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
