package routes

import (
	"github.com/gin-gonic/gin"
	authTransport "github.com/richardktran/MyBlogBE/modules/auth/transports"
)

func V1(router *gin.RouterGroup) {
	router.POST("/login", authTransport.Login)
}
