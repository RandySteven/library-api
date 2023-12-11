package main

import (
	"context"
	"errors"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func AuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if info.FullMethod != "/library.AuthService/Login" {
		return handler(ctx, req)
	}

	if !ok {
		return nil, errors.New("missing token")
	}

	tokenHeader := strings.TrimPrefix(md["Authorization"][0], "Bearer ")

	token := utils.ValidateToken(tokenHeader)
	if token == nil {
		return nil, errors.New("error generate token jwt")
	}

	ctx = context.WithValue(ctx, "id", token.ID)
	ctx = context.WithValue(ctx, "email", token.Email)

	return handler(ctx, req)
}

func ErrorInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	res, err := handler(ctx, req)

	return res, err
}
