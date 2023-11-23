package main

import (
	"os"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
)

func InitConfig() *models.Config {
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	return models.NewConfig(dbHost, dbPort, dbUser, dbPass, dbName)
}

func AppPort() string {
	return os.Getenv("APP_PORT")
}
