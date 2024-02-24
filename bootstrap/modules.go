package bootstrap

import (
	"github.com/richardktran/MyBlogBE/app/handlers"
	repositories "github.com/richardktran/MyBlogBE/app/respositories"
	"github.com/richardktran/MyBlogBE/app/services"
	"github.com/richardktran/MyBlogBE/pkg/router"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	handlers.Module,
	router.Module,
	services.Module,
	repositories.Module,
)
