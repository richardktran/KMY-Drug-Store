package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/richardktran/MyBlogBE/routes"
)

func loadVersionRoutes(router *gin.RouterGroup, version string) {
	switch version {
	case "v1":
		routes.V1(router)
	default:
		panic(fmt.Sprintf("Version %s is not supported", version))
	}
}

func LoadAPIRouter(router *gin.Engine, version string) {
	path := fmt.Sprintf("%s/%s", "api", version)
	api := router.Group(path)
	{
		loadVersionRoutes(api, version)
	}
}
