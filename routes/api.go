package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/richardktran/KMY-Drug-Store/app/controllers"
)

type ApiV1Route struct {
	orderController controllers.OrderController
}

func NewApiV1Route(orderController controllers.OrderController) ApiV1Route {
	return ApiV1Route{
		orderController: orderController,
	}
}

func (r ApiV1Route) Setup(router *gin.Engine) {
	api := router.Group("/api/v1/")
	{
		todo := api.Group("/orders")
		{
			todo.POST("/", r.orderController.StoreOrder())
		}
	}
}
