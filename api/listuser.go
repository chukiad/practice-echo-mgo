package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func ListUser(c echo.Context) error {
	var u []User

	return c.JSON(http.StatusOK, u)
}
