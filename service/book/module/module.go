package module

import (
	"go.uber.org/fx"

	bookHTTP "exchequer/service/book/delivery/http"
	bookRepo "exchequer/service/book/repository/postgres"
	bookUsecase "exchequer/service/book/usecase"
)

var Module = fx.Options(
	fx.Provide(
		bookHTTP.New,
		bookUsecase.New,
		bookRepo.New,
	),
)
