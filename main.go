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
	v1.POST("/login", handlers.AuthHandler.LoginUser)
	v1.POST("/register", handlers.AuthHandler.RegisterUser)
	v1.POST("/logout", handlers.AuthHandler.LogoutUser)
	handlers.InitRouter(v1)

	srv := http.Server{
		Addr:    ":" + AppPort(),
		Handler: r,
	}
	srv.ListenAndServe()
}
