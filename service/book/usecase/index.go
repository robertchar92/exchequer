package usecase

import (
	"exchequer/models"
	"exchequer/utils/request_util"
	response_util "exchequer/utils/response_utils"
)

func (u *Usecase) Index(paginationConfig request_util.PaginationConfig) ([]models.Book, response_util.PaginationMeta, error) {
	meta := response_util.PaginationMeta{
		Offset: paginationConfig.Offset(),
		Limit:  paginationConfig.Limit(),
		Total:  0,
	}

	books, err := u.bookRepo.FindAll(paginationConfig)
	if err != nil {
		return nil, meta, err
	}

	return books, meta, nil
}
