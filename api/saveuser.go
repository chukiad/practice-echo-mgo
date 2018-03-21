package api

import (
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/labstack/echo"
)

// SaveUser : func to new user in userdata collection
func SaveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, RspMsg{err.Error()})
	}
	if u.Name == "" || u.Email == "" {
		return c.JSON(http.StatusBadRequest, RspMsg{"Invalid Parameter"})
	}
	u.ID = bson.NewObjectId()
	u.Time = time.Now()

	err := Dbsession.DB("").C("UserData").Insert(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, RspMsg{err.Error()})
	}
	return c.JSON(http.StatusCreated, u)
}
