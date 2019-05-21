package api

import "github.com/labstack/echo"

// RegisterRoutes API EndPoint
func RegisterRoutes(g *echo.Group) {
	role := g.Group("role")
	role.GET("/:role_id", getRole)
	role.GET("", getRoles)
	role.POST("", createRole)
}
