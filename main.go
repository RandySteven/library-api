package main

import (
	"log"
	"net/http"
	"os"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/configs"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/server"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env got")
	}
	r := gin.Default()

	//TO DO : take out from main
	appPort := os.Getenv("APP_PORT")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	config := models.NewConfig(dbHost, dbPort, dbUser, dbPass, dbName)

	//TO DO : take out from main
	repository, err := configs.NewRepository(config)
	if err != nil {
		log.Println(err)
		return
	}
	defer repository.Close()
	err = repository.Automigrate()
	if err != nil {
		return
	}

	handlers, err := server.NewHandlers(*repository)
	if err != nil {
		return
	}

	v1 := r.Group("/v1")
	handlers.InitRouter(v1)

	srv := http.Server{
		Addr:    ":" + appPort,
		Handler: r,
	}
	srv.ListenAndServe()
}
