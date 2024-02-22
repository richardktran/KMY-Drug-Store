package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/richardktran/MyBlogBE/pkg/app"
	"github.com/richardktran/MyBlogBE/routes"
)

var routerInstance *gin.Engine

func initRouter() *gin.Engine {
	if app.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	registerAPIRoutes(r, "v1")
	registerWebRoutes(r)

	return r
}

func GetRouter() *gin.Engine {
	if routerInstance == nil {
		routerInstance = initRouter()
	}

	return routerInstance
}

func registerWebRoutes(router *gin.Engine) {
	web := router.Group("/")
	{
		routes.Web(web)
	}
}

func registerAPIRoutes(router *gin.Engine, version string) {
	path := fmt.Sprintf("%s/%s", "api", version)
	api := router.Group(path)
	{
		loadVersionRoutes(api, version)
	}
}

func loadVersionRoutes(router *gin.RouterGroup, version string) {
	switch version {
	case "v1":
		routes.V1(router)
	default:
		panic(fmt.Sprintf("Version %s is not supported", version))
	}
}
