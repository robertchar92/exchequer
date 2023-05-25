package postgres

import (
	"log"

	"gorm.io/gorm"

	"exchequer/models"
	"exchequer/utils/errors"
)

func (r *Repository) FindByEmail(email string) (*models.User, error) {
	model := models.User{}

	err := r.db.Where("email = ?", email).Set("cached", false).First(&model).Error

	if err == gorm.ErrRecordNotFound {
		return nil, errors.ErrNotFound
	}

	if err != nil {
		log.Println("error-find-user-by-email:", err)
		return nil, errors.CustomWrap(err)
	}

	return &model, nil
}
