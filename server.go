package main

import (
	"log"
	"practice-echo-mgo/api"

	"gopkg.in/mgo.v2"

	"github.com/labstack/echo"
)

func main() {
	var err error

	api.Dbsession, err = mgo.Dial("localhost/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer api.Dbsession.Close()

	e := echo.New()

	e.GET("/users", api.ListUser)
	e.GET("/users/:id", api.GetUser)
	e.POST("/users", api.SaveUser)
	e.PUT("/users/:id", api.UpdateUser)
	e.DELETE("/users/:id", api.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
