package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetUser(c echo.Context) error {
	id := c.Param("id")
	u := User{id, "john", "john@gmail.com"}
	return c.JSON(http.StatusOK, u)
}
