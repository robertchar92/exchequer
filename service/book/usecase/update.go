package usecase

import (
	"exchequer/models"
	"exchequer/service/book/delivery/http/request"
	"exchequer/utils/errors"

	"github.com/jinzhu/copier"
)

func (u *Usecase) Update(bookID string, userID string, request request.BookUpdateRequest) (*models.Book, error) {
	bookM, err := u.bookRepo.FindByID(bookID)
	if err != nil {
		err := errors.ErrNotFound
		err.Message = "Book not found!"
		return nil, err
	}

	if userID != bookM.UserID {
		err := errors.ErrUnauthorized
		err.Message = "You are not the owner of the book!"
		return nil, err
	}

	_ = copier.Copy(bookM, &request)

	err = u.bookRepo.Update(bookM, nil)
	if err != nil {
		err := errors.ErrUnprocessableEntity
		err.Message = "Failed to update book"
		return nil, err
	}

	return bookM, nil
}
