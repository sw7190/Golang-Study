package main

import (
	"github.com/labstack/echo"

	"echo-study/api"
	"echo-study/db"
)

func init() {
	db.InitDataBase()
}

func main() {
	defer func() {
		if err := recover(); err != nil {

		}
	}()

	e := echo.New()
	// g := e.Group("/")
	// api.RegisterRoutes(g)

	e.GET("/list", api.GetList)

	if err := e.Start(":5000"); nil != err {
	}
}
