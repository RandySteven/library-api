package interfaces

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/pb"
)

type (
	BorrowGrpcRepository interface {
		CreateBorrowRepo(ctx context.Context, req *pb.BorrowRequest) (*pb.BorrowBookResponse, error)
	}

	BorrowGrpcUsecase interface {
		CreateBorrowRecord(ctx context.Context, req *pb.BorrowRequest) (*pb.BorrowBookResponse, error)
	}
)
