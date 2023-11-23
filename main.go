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

	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	config := models.NewConfig(dbHost, dbPort, dbUser, dbPass, dbName)

	service, err := configs.NewRepository(config)
	if err != nil {
		log.Println(err)
		return
	}
	defer service.Close()
	err = service.Automigrate()
	if err != nil {
		return
	}

	handlers, err := server.NewHandlers(*service)
	if err != nil {
		return
	}

	v1 := r.Group("/v1")
	handlers.InitRouter(v1)

	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	srv.ListenAndServe()
	os.Exit(0)
}
