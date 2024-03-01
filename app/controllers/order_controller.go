package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/richardktran/KMY-Drug-Store/app/models"
	"github.com/richardktran/KMY-Drug-Store/app/services/contracts"
	"github.com/richardktran/KMY-Drug-Store/pkg/app"
)

type OrderController struct {
	orderService contracts.IOrderService
}

func NewOrderController(
	orderService contracts.IOrderService,
) OrderController {
	return OrderController{
		orderService: orderService,
	}
}

func (o OrderController) StoreOrder() func(c *gin.Context) {
	return func(c *gin.Context) {
		var orderRequest models.OrderCreation

		if err := c.ShouldBind(&orderRequest); err != nil {
			app.ResponseBadRequest(
				app.ThrowBadRequestError(err, "invalid_request"),
			).Context(c)

			return
		}

		order, exception := o.orderService.StoreOrder(&orderRequest)

		if exception != nil {
			app.ResponseBadRequest(exception).Context(c)

			return
		}

		app.ResponseSuccess(order).Context(c)
	}
}
