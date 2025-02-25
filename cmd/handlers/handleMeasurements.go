package handlers

import (
	"echo-rest/cmd/models"
	"echo-rest/cmd/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateMeasurements(c echo.Context) error {
	measurement := models.Measurements{}
	c.Bind(&measurement)
	newMeasurement, err := repositories.CreateMeasurements(measurement)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, newMeasurement)
}
