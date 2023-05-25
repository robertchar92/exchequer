package module

import (
	"go.uber.org/fx"

	userHTTP "exchequer/service/user/delivery/http"
	userRepo "exchequer/service/user/repository/postgres"
	userUsecase "exchequer/service/user/usecase"
)

var Module = fx.Options(
	fx.Provide(
		userHTTP.New,
		userUsecase.New,
		userRepo.New,
	),
)
