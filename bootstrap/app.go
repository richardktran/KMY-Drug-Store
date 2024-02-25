package bootstrap

import (
	"log"

	"github.com/richardktran/MyBlogBE/app/handlers"
	repositories "github.com/richardktran/MyBlogBE/app/respositories"
	"github.com/richardktran/MyBlogBE/app/services"
	"github.com/richardktran/MyBlogBE/pkg/database"
	"github.com/richardktran/MyBlogBE/pkg/middleware"
	"github.com/richardktran/MyBlogBE/pkg/router"
	"go.uber.org/fx"
)

var Module = fx.Options(
	handlers.Module,
	router.Module,
	services.Module,
	repositories.Module,
	fx.Provide(router.NewRoutes),
	fx.Invoke(RunApp),
)

func RunApp(router router.Routes) {
	r := router.Setup()
	// Define middleware here
	r.Use(middleware.Recovery())

	r = router.RegisterRoutes(r)

	if r == nil {
		log.Fatal("Failed to initialize router")
	}

	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.Run("localhost:3000")

	defer database.CloseDB()
}
