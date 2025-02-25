package repositories

import (
	"echo-rest/cmd/models"
	"echo-rest/cmd/storage"
	"time"
)

func CreateMeasurements(measurements models.Measurements) (models.Measurements, error) {
	db := storage.GetDB()
	sqlStatement := `INSERT INTO measurements (user_id, weight, height, body_fat, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRow(sqlStatement, measurements.UserId, measurements.Weight, measurements.Height, measurements.BodyFat, time.Now()).Scan(&measurements.Id)
	if err != nil {
		return measurements, err
	}
	return measurements, nil
}
