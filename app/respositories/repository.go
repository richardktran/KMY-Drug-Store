package repositories

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewOrderRepository),
	fx.Provide(NewUserRepository),
	fx.Provide(NewProductRepository),
)
