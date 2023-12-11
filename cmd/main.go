package main

import (
	"log"
	"net"
	"os"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/configs"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	handler_grpc "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler/grpc"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/pb"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/usecase"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func InitConfig() *models.Config {
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	return models.NewConfig(dbHost, dbPort, dbUser, dbPass, dbName)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env got")
	}

	config := InitConfig()

	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Println("err : ", err)
		return
	}

	opt := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(),
	}

	server := grpc.NewServer(opt...)

	repository, err := configs.NewRepository(config)
	if err != nil {
		log.Println(err)
		return
	}

	authUsecase := usecase.NewAuthUseCase(repository.AuthRepository)
	auth := handler_grpc.NewAuthHandler(authUsecase)

	pb.RegisterAuthServiceServer(server, auth)

	if err = server.Serve(listener); err != nil {
		log.Fatal("err : ", err)
	}
	log.Println("server running at :50053")

}
