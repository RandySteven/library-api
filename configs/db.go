package configs

import (
	"fmt"
	"log"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository"
	_ "github.com/jackc/pgx/v4/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	AuthorRepository interfaces.AuthorRepository
	BookRepository   interfaces.BookRepository
	BorrowRepository interfaces.BorrowRepository
	UserRepository   interfaces.UserRepository
	AuthRepository   interfaces.AuthRepository
	db               *gorm.DB
}

func NewRepository(config *models.Config) (*Repository, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost,
		config.DbPort,
		config.DbUser,
		config.DbPass,
		config.DbName,
	)
	log.Println(conn)
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println("DB error : db.go : ", err)
		return nil, err
	}
	return &Repository{
		AuthorRepository: repository.NewAuthorRepository(db),
		BookRepository:   repository.NewBookRepository(db),
		UserRepository:   repository.NewUserRepository(db),
		BorrowRepository: repository.NewBorrowRepository(db),
		AuthRepository:   repository.NewAuthRepository(db),
		db:               db,
	}, nil
}

func (r *Repository) Automigrate() error {
	return r.db.AutoMigrate(
		&models.Author{},
		&models.User{},
		&models.Book{},
		&models.Borrow{},
	)
}

func (r *Repository) TrxBegin() *gorm.DB {
	return r.db.Begin()
}
