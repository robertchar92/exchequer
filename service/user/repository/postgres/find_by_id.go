package postgres

import (
	"log"

	"gorm.io/gorm"

	"exchequer/models"
	"exchequer/utils/errors"
)

func (r *Repository) FindByID(userID string, unscoped ...bool) (*models.User, error) {
	model := models.User{}

	q := r.db
	if len(unscoped) > 0 && unscoped[0] {
		q = q.Unscoped()
	}
	err := q.Where("id = ?", userID).First(&model).Error

	if err == gorm.ErrRecordNotFound {
		return nil, errors.ErrNotFound
	}

	if err != nil {
		log.Println("error-find-user-by-id:", err)
		return nil, errors.CustomWrap(err)
	}

	return &model, nil
}
