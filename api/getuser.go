package api

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/labstack/echo"
)

// GetUser : get specific id userdata
func GetUser(c echo.Context) error {
	id := c.Param("id")
	u := new(User)

	if !bson.IsObjectIdHex(id) {
		return c.JSON(http.StatusBadRequest, RspMsg{"Invalid ObjectIdHex"})
	}
	err := Dbsession.DB("").C("UserData").FindId(bson.ObjectIdHex(id)).One(&u)
	if err != nil {
		return c.JSON(http.StatusNotFound, RspMsg{err.Error()})
	}
	return c.JSON(http.StatusOK, u)
}
