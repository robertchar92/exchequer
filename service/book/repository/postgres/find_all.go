package postgres

import (
	"exchequer/models"
	"exchequer/utils/errors"
	"exchequer/utils/request_util"
	"log"
)

func (r *Repository) FindAll(config request_util.PaginationConfig) ([]models.Book, error) {
	results := make([]models.Book, 0)

	err := r.db.
		Scopes(config.MetaScopes()...).
		Scopes(config.Scopes()...).
		Find(&results).Error
	if err != nil {
		log.Println("error-find-book:", err)
		return nil, errors.CustomWrap(err)
	}

	return results, nil
}
