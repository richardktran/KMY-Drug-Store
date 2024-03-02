package controllers

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewOrderController),
	fx.Provide(NewUserController),
)
