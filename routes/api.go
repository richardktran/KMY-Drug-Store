package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/richardktran/KMY-Drug-Store/app/controllers"
)

type ApiV1Route struct {
	orderController   controllers.OrderController
	userController    controllers.UserController
	productController controllers.ProductController
	reportController  controllers.ReportController
}

func NewApiV1Route(
	orderController controllers.OrderController,
	userController controllers.UserController,
	productController controllers.ProductController,
	reportController controllers.ReportController,
) ApiV1Route {
	return ApiV1Route{
		orderController:   orderController,
		userController:    userController,
		productController: productController,
		reportController:  reportController,
	}
}

func (r ApiV1Route) Setup(router *gin.Engine) {
	api := router.Group("/api/v1/")
	{
		admin := api.Group("/admin")
		{
			admin.GET("/orders", r.orderController.GetOrders())
		}

		orders := api.Group("/orders")
		{
			orders.GET("", r.orderController.GetOrdersByPhoneNumber())
			orders.POST("", r.orderController.StoreOrder())
		}

		user := api.Group("/users")
		{
			user.GET("", r.userController.GetUserByPhone())
			user.GET("/list", r.userController.GetUserList())
		}

		product := api.Group("/products")
		{
			product.GET("", r.productController.GetProduct())
		}

		report := api.Group("/reports")
		{
			report.GET("/revenues", r.reportController.GetRevenues())
		}
	}
}
