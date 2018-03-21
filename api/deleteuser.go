package api

import (
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	u := new(User)

	if !bson.IsObjectIdHex(id) {
		return c.JSON(http.StatusBadRequest, RspMsg{"Invalid ObjectIdHex"})
	}

	uc := Dbsession.DB("").C("UserData")
	var err error
	err = uc.FindId(bson.ObjectIdHex(id)).One(&u)
	if err != nil {
		return c.JSON(http.StatusNotFound, RspMsg{err.Error()})
	}
	err = uc.RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, RspMsg{err.Error()})
	}
	return c.JSON(http.StatusOK, u)
}
