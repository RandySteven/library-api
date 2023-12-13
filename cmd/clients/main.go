package main

import (
	"context"
	"log"

	handler_grpc "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler/grpc"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/pb"
	repository "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository/grpc"
	usecase "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/usecase/grpc"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("10.20.191.19:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	defer conn.Close()

	client := pb.NewBorrowBookServiceClient(conn)

	req := &pb.BorrowRequest{
		BookId: 3,
	}

	borrowRepo := repository.NewBorrowGrpcRepository(conn, client)
	borrowUsecase := usecase.NewBorrowGrpcUsecase(borrowRepo)
	borrowHandler := handler_grpc.NewBorrowGrpcHandler(borrowUsecase)

	res, err := borrowHandler.CreateBorrowBook(context.Background(), req)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	log.Println(res)
}
