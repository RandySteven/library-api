package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env got")
	}
	r := gin.Default()
	r.ContextWithFallback = true

	handlers := InitHandlers()

	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.ErrorMiddleware())
	v1 := r.Group("/v1")
	v1.POST("/login", handlers.AuthHandler.LoginUser)
	v1.POST("/register", handlers.AuthHandler.RegisterUser)
	v1.POST("/logout", handlers.AuthHandler.LogoutUser)
	handlers.InitRouter(v1)

	srv := http.Server{
		Addr:    ":" + AppPort(),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
