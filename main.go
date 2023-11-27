package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env got")
	}
	r := gin.Default()

	handlers := InitHandlers()

	v1 := r.Group("/v1")
	handlers.InitRouter(v1)

	srv := http.Server{
		Addr:    ":" + AppPort(),
		Handler: r,
	}
	srv.ListenAndServe()
}
