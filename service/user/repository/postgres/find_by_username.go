package postgres

import (
	"log"

	"gorm.io/gorm"

	"exchequer/models"
	"exchequer/utils/errors"
)

func (r *Repository) FindByUsername(username string, unscoped ...bool) (*models.User, error) {
	model := models.User{}

	q := r.db
	if len(unscoped) > 0 && unscoped[0] {
		q = q.Unscoped()
	}
	err := q.Where("username = ?", username).Set("cached", false).First(&model).Error

	if err == gorm.ErrRecordNotFound {
		return nil, errors.ErrNotFound
	}

	if err != nil {
		log.Println("error-find-user-by-username:", err)
		return nil, errors.CustomWrap(err)
	}

	return &model, nil
}
