package repository

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type borrowGrpcRepository struct {
	conn   *grpc.ClientConn
	borrow pb.BorrowBookServiceClient
}

// CreateBorrowRepo implements interfaces.BorrowGrpcRepository.
func (repo *borrowGrpcRepository) CreateBorrowRepo(ctx context.Context, req *pb.BorrowRequest) (*pb.BorrowBookResponse, error) {
	md := metadata.Pairs("authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDI5NzUwNzAsImlzcyI6IkxpYnJhcnkgQVBJIiwidXNlcl9pZCI6NjgsIm5hbWUiOiJhbGRpIiwiZW1haWwiOiJhbGR5c3AzM0BnbWFpbC5jb20iLCJwaG9uZSI6IjA4NzY4ODAxMjM3In0.a0ZqAWQ51QQqwWMozT3RS5a3y0qzNjFHaGMJ1Xwa7GI")
	ctx = metadata.NewOutgoingContext(ctx, md)

	return repo.borrow.CreateBorrowBook(ctx, req)
}

func NewBorrowGrpcRepository(conn *grpc.ClientConn, borrow pb.BorrowBookServiceClient) *borrowGrpcRepository {
	return &borrowGrpcRepository{
		conn:   conn,
		borrow: borrow,
	}
}

var _ interfaces.BorrowGrpcRepository = &borrowGrpcRepository{}
