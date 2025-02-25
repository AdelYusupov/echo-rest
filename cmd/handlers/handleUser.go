package handlers

import (
	"echo-rest/cmd/models"
	"echo-rest/cmd/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	newUser, err := repositories.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, newUser)
}
