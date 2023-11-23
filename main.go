package main

import (
	"log"
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/configs"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/server"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env got")
	}
	r := gin.Default()

	config := InitConfig()

	repository, err := configs.NewRepository(config)
	if err != nil {
		log.Println(err)
		return
	}

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
		Addr:    ":" + AppPort(),
		Handler: r,
	}
	srv.ListenAndServe()
}
