package main

import (
	"practice-echo-mgo/api"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/users", api.ListUser)
	e.GET("/users/:id", api.GetUser)
	e.POST("/users", api.SaveUser)

	e.Logger.Fatal(e.Start(":1323"))
}
