package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func SaveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}
