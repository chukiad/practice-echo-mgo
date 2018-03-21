package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

// UpdateUser : update userdata by id
func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	u := new(UserForUpdate)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, RspMsg{err.Error()})
	}
	if !bson.IsObjectIdHex(id) {
		return c.JSON(http.StatusBadRequest, RspMsg{"Invalid ObjectIdHex"})
	}
	if u.Name == "" || u.Email == "" {
		return c.JSON(http.StatusBadRequest, RspMsg{"Invalid Parameter"})
	}
	u.Time = time.Now()
	err := Dbsession.DB("").C("UserData").UpdateId(bson.ObjectIdHex(id), u)
	if err != nil {
		return c.JSON(http.StatusNotFound, RspMsg{err.Error()})
	}
	return c.JSON(http.StatusOK, u)
}
