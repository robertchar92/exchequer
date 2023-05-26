package postgres

import (
	"log"

	"exchequer/models"
	"exchequer/utils/errors"
	"exchequer/utils/request_util"
)

func (r *Repository) Count(config request_util.PaginationConfig) (int64, error) {
	var count int64

	err := r.db.
		Model(&models.Book{}).
		Scopes(config.Scopes()...).
		Count(&count).Error
	if err != nil {
		log.Println("error-count-book:", err)
		return 0, errors.CustomWrap(err)
	}

	return count, nil
}
