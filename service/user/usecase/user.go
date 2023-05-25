package usecase

import "exchequer/service/user"

type Usecase struct {
	userRepo user.Repository
}

func New(
	userRepo user.Repository,
) user.Usecase {
	return &Usecase{
		userRepo: userRepo,
	}
}
