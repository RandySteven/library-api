package configs

import "gorm.io/gorm"

type Transaction struct {
	repo *Repository
}

func NewTransaction(repo *Repository) *Transaction {
	return &Transaction{repo}
}

func (t *Transaction) Begin() *gorm.DB {
	return t.repo.db.Begin()
}

func (t *Transaction) Rollback() *gorm.DB {
	return t.Begin().Rollback()
}

func (t *Transaction) Commit() *gorm.DB {
	return t.Begin().Commit()
}
