package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/richardktran/KMY-Drug-Store/app/controllers"
)

type ApiV1Route struct {
	orderController controllers.OrderController
	userController  controllers.UserController
}

func NewApiV1Route(
	orderController controllers.OrderController,
	userController controllers.UserController,
) ApiV1Route {
	return ApiV1Route{
		orderController: orderController,
		userController:  userController,
	}
}

func (r ApiV1Route) Setup(router *gin.Engine) {
	api := router.Group("/api/v1/")
	{
		todo := api.Group("/orders")
		{
			todo.POST("/", r.orderController.StoreOrder())
		}

		user := api.Group("/users")
		{
			user.GET("/:phone_number", r.userController.GetUserByPhone())
		}
	}
}
