package main

import (
	"fmt"

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
			fmt.Println("main1")
		}
	}()

	e := echo.New()
	// g := e.Group("/")
	// api.RegisterRoutes(g)

	e.GET("/list", api.GetList)

	if err := e.Start(":5000"); nil != err {
		fmt.Println("main2")
	}
}
