package api

import (
	"net/http"

	"github.com/labstack/echo"
)

// ListUser : list all userdata
func ListUser(c echo.Context) error {
	var u []User
	err := Dbsession.DB("").C("UserData").Find(nil).All(&u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, RspMsg{err.Error()})
	}
	return c.JSON(http.StatusOK, u)
}
