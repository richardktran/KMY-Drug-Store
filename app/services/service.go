package services

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewUserService),
	fx.Provide(NewOrderService),
	fx.Provide(NewProductService),
)
