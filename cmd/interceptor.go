package main

import (
	"context"
	"errors"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	res, err := handler(ctx, req)

	var errBadRequest *apperror.ErrBadRequest
	var errNoDuplication *apperror.ErrNoDuplication
	var errBookIdNotFound *apperror.ErrBookIdNotFound
	var errUserIdNotFound *apperror.ErrUserIdNotFound
	var errBookQuantityZero *apperror.ErrBookQuantityZero
	var errBorrowStatusAlreadyReturned *apperror.ErrBorrowStatusAlreadyReturned
	var errUnauthorized *apperror.ErrUnauthorized
	var errBorrowRecordNotFound *apperror.ErrBorrowRecordNotFound
	var errPermissionDenied *apperror.ErrPermissionDenied
	var errPasswordTooLong *apperror.ErrPasswordTooLong

	if err != nil {
		switch {
		case errors.As(err, &errPasswordTooLong):
			return res, status.Error(codes.InvalidArgument, err.Error())
		case errors.As(err, &errBadRequest):
			return res, status.Error(codes.InvalidArgument, err.Error())
		case errors.As(err, &errNoDuplication):
			return res, status.Error(codes.InvalidArgument, err.Error())
		case errors.As(err, &errBookIdNotFound):
			return res, status.Error(codes.InvalidArgument, err.Error())
		case errors.As(err, &errUserIdNotFound):
			return res, status.Error(codes.InvalidArgument, err.Error())
		case errors.As(err, &errBookQuantityZero):
			return res, status.Error(codes.InvalidArgument, err.Error())
		case errors.As(err, &errBorrowStatusAlreadyReturned):
			return res, status.Error(codes.InvalidArgument, err.Error())
		case errors.As(err, &errUnauthorized):
			return res, status.Error(codes.Unauthenticated, err.Error())
		case errors.As(err, &errBorrowRecordNotFound):
			return res, status.Error(codes.InvalidArgument, err.Error())
		case errors.As(err, &errPermissionDenied):
			return res, status.Error(codes.InvalidArgument, err.Error())
		default:
			return res, status.Error(codes.Unknown, err.Error())
		}
	}

	return res, err
}

func LogInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	var log = logrus.New()

	res, err := handler(ctx, req)

	log.Info("Print")

	return res, err
}
