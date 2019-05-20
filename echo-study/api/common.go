package api

import "github.com/labstack/echo"

func RegisterRoutes(g *echo.Group) {
	g.GET("/list", getList)
}
