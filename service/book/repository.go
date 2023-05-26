package book

import (
	"exchequer/models"
	"exchequer/utils/request_util"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(config request_util.PaginationConfig) ([]models.Book, error)
	FindByID(bookID string) (*models.Book, error)
	Count(config request_util.PaginationConfig) (int64, error)
	Insert(book *models.Book, tx *gorm.DB) error
	Update(book *models.Book, tx *gorm.DB) error
}
