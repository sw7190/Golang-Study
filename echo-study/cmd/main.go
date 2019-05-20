package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// git user test
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Wrold!")
	})

	e.Start(":5000")
}
