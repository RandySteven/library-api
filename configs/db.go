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
	BookRepository interfaces.BookRepository
	db             *gorm.DB
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
	// defer db.Close()
	return &Repository{
		BookRepository: repository.NewBookRepository(db),
		db:             db,
	}, nil
}

func (repo *Repository) Close() <-chan struct{} {
	return repo.db.Statement.Context.Done()
}

func (r *Repository) Automigrate() error {
	return r.db.AutoMigrate(
		&models.Book{},
	)
}
