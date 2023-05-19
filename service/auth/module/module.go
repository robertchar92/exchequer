package module

import (
	"go.uber.org/fx"

	authHTTP "exchequer/service/auth/delivery/http"
	authUsecase "exchequer/service/auth/usecase"
)

var Module = fx.Options(
	fx.Provide(
		authHTTP.New,
		authUsecase.New,
	),
)
