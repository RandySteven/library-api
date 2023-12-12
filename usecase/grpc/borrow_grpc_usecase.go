package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/pb"
)

type borrowGrpcUsecase struct {
	repo interfaces.BorrowGrpcRepository
}

// CreateBorrowRecord implements interfaces.BorrowGrpcUsecase.
func (usecase *borrowGrpcUsecase) CreateBorrowRecord(ctx context.Context, req *pb.BorrowRequest) (*pb.BorrowBookResponse, error) {
	return usecase.repo.CreateBorrowRepo(ctx, req)
}

func NewBorrowGrpcUsecase(repo interfaces.BorrowGrpcRepository) *borrowGrpcUsecase {
	return &borrowGrpcUsecase{repo: repo}
}

var _ interfaces.BorrowGrpcUsecase = &borrowGrpcUsecase{}
