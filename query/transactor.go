package query

import (
	"context"

	"gorm.io/gorm"
)

type transactor struct {
	db *gorm.DB
}

func NewTransactor(db *gorm.DB) *transactor {
	return &transactor{db: db}
}

func (t *transactor) Transaction(ctx context.Context, txFunc func(txCtx context.Context) error) error {
	err := t.db.Transaction(func(tx *gorm.DB) error {
		txCtx := context.WithValue(ctx, "tx", tx)
		err := txFunc(txCtx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
