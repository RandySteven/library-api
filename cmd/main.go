package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/configs"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	handler_grpc "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler/grpc"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/pb"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/usecase"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func InitConfig() *models.Config {
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	return models.NewConfig(dbHost, dbPort, dbUser, dbPass, dbName)
}

func valid(tokenString string) *configs.JWTClaim {

	if tokenString == "" {
		return nil
	}

	claims := &configs.JWTClaim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return configs.JWT_KEY, nil
	})

	if err != nil || !token.Valid {
		log.Println("err token : ", err)
		log.Println("token valid : ", token.Valid)
		return nil
	}

	return claims
}

func AuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if info.FullMethod == "/library.AuthService/Login" {
		return handler(ctx, req)
	}

	if !ok {
		return nil, errors.New("missing token")
	}

	tokenHeader := md["authorization"][0]
	token := valid(tokenHeader)
	if token == nil {
		return nil, errors.New("error get token")
	}

	ctx = context.WithValue(ctx, "id", token.ID)
	ctx = context.WithValue(ctx, "name", token.Name)
	ctx = context.WithValue(ctx, "email", token.Email)

	return handler(ctx, req)
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
		grpc.ChainUnaryInterceptor(AuthInterceptor),
	}

	server := grpc.NewServer(opt...)

	repository, err := configs.NewRepository(config)
	if err != nil {
		log.Println(err)
		return
	}

	authUsecase := usecase.NewAuthUseCase(repository.AuthRepository)
	userUsecase := usecase.NewUserUseCase(repository.UserRepository)
	bookUsecase := usecase.NewBookUseCase(repository.BookRepository, repository.AuthorRepository)
	auth := handler_grpc.NewAuthHandler(authUsecase)
	user := handler_grpc.NewUserHandler(userUsecase)
	book := handler_grpc.NewBookHandler(bookUsecase)

	pb.RegisterAuthServiceServer(server, auth)
	pb.RegisterUserServiceServer(server, user)
	pb.RegisterBookServiceServer(server, book)

	if err = server.Serve(listener); err != nil {
		log.Fatal("err : ", err)
	}
	log.Println("server running at :50053")

}
