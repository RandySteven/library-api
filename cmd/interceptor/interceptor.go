package interceptor

import (
	"context"
	"errors"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/logger"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if info.FullMethod == "/library.AuthService/Login" {
		return handler(ctx, req)
	}

	if !ok {
		return nil, status.Error(codes.Unauthenticated, "unauthroized")
	}

	tokenHeader := md["authorization"][0]
	token := utils.ValidateToken(tokenHeader)
	if token == nil {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	ctx = context.WithValue(ctx, "id", token.ID)
	ctx = context.WithValue(ctx, "name", token.Name)
	ctx = context.WithValue(ctx, "email", token.Email)

	return handler(ctx, req)
}

func ErrorInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	res, err := handler(ctx, req)

	var (
		errBadRequest                  *apperror.ErrBadRequest
		errNoDuplication               *apperror.ErrNoDuplication
		errBookIdNotFound              *apperror.ErrBookIdNotFound
		errUserIdNotFound              *apperror.ErrUserIdNotFound
		errBookQuantityZero            *apperror.ErrBookQuantityZero
		errBorrowStatusAlreadyReturned *apperror.ErrBorrowStatusAlreadyReturned
		errUnauthorized                *apperror.ErrUnauthorized
		errBorrowRecordNotFound        *apperror.ErrBorrowRecordNotFound
		errPermissionDenied            *apperror.ErrPermissionDenied
		errPasswordTooLong             *apperror.ErrPasswordTooLong
	)

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
			return res, status.Error(codes.PermissionDenied, err.Error())
		default:
			return res, status.Error(codes.Unknown, err.Error())
		}
	}

	return res, err
}

func LogInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	res, err := handler(ctx, req)

	log := logger.NewLog()

	start := time.Now()
	args := map[string]interface{}{
		"status":  status.Code(err),
		"latency": time.Since(start),
		"path":    info.FullMethod,
	}

	if err != nil {
		args["err"] = err.Error()
		log.Error(args)
	} else {
		args["req"] = utils.Encode(req)
		log.Info(args)
	}

	return res, err
}
