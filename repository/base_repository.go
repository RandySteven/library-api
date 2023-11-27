package repository

import (
	"context"

	"gorm.io/gorm"
)

type BaseRepository interface {
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func extraTx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("tx").(*gorm.DB); ok {
		return tx
	}
	return nil
}

func (r *baseRepository[T]) conn(ctx context.Context) *gorm.DB {
	tx := extraTx(ctx)
	if tx != nil {
		return tx
	}
	return r.db
}
