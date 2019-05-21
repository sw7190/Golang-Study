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
			fmt.Printf("server panic; err: %v \n", err)
		}
	}()

	e := echo.New()
	g := e.Group("/")
	api.RegisterRoutes(g)

	if err := e.Start(":5000"); nil != err {
		fmt.Printf("server panic; err: %v \n", err)
	}
}
