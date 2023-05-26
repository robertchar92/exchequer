package usecase

import (
	"exchequer/service/book"
	"exchequer/service/user"
)

type Usecase struct {
	userRepo user.Repository
	bookRepo book.Repository
}

func New(
	userRepo user.Repository,
	bookRepo book.Repository,
) book.Usecase {
	return &Usecase{
		userRepo: userRepo,
		bookRepo: bookRepo,
	}
}
