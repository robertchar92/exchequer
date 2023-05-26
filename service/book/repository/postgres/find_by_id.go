package postgres

import (
	"log"

	"gorm.io/gorm"

	"exchequer/models"
	"exchequer/utils/errors"
)

func (r *Repository) FindByID(bookID string) (*models.Book, error) {
	model := models.Book{}

	err := r.db.Where("id = ?", bookID).First(&model).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.ErrNotFound
	}

	if err != nil {
		log.Println("error-find-book-by-id:", err)
		return nil, errors.CustomWrap(err)
	}

	return &model, nil
}
