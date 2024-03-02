package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/richardktran/KMY-Drug-Store/app/models"
	"github.com/richardktran/KMY-Drug-Store/app/services/contracts"
	"github.com/richardktran/KMY-Drug-Store/pkg/app"
)

type OrderController struct {
	orderService contracts.IOrderService
	userService  contracts.IUserService
}

func NewOrderController(
	orderService contracts.IOrderService,
	userService contracts.IUserService,
) OrderController {
	return OrderController{
		orderService: orderService,
		userService:  userService,
	}
}

func (ctl *OrderController) GetOrders() func(*gin.Context) {
	return func(c *gin.Context) {
		phoneNumber := c.DefaultQuery("phone_number", "")

		if phoneNumber == "" {
			app.ResponseBadRequest(
				app.ThrowBadRequestError(errors.New("phone_number_is_required"), "phone_number_is_required"),
			).Context(c)

			return
		}

		user, err := ctl.userService.GetUserByPhoneNumber(phoneNumber)

		if err != nil {
			app.ResponseBadRequest(
				app.ThrowNotFoundError(errors.New("user_not_found"), "user_not_found"),
			).Context(c)

			return
		}

		orders, meta, err := ctl.orderService.GetAllOrders(
			map[string]interface{}{"user_id": user.ID},
			false,
		)

		if err != nil {
			app.ResponseNotFound(
				app.ThrowNotFoundError(err, "orders_not_found"),
			).Context(c)

			return
		}

		app.ResponseSuccessWithMetaData(orders, meta).Context(c)
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

		if orderRequest.Amount == 0 {
			app.ResponseBadRequest(
				app.ThrowBadRequestError(nil, "amount_is_required"),
			).Context(c)

			return
		}

		if orderRequest.PhoneNumber == "" {
			app.ResponseBadRequest(
				app.ThrowBadRequestError(nil, "phone_number_is_required"),
			).Context(c)

			return
		}

		_, exception := o.userService.GetUserByPhoneNumber(orderRequest.PhoneNumber)

		if exception != nil {
			if orderRequest.FullName == "" {
				app.ResponseBadRequest(
					app.ThrowBadRequestError(nil, "full_name_is_required"),
				).Context(c)

				return
			}
		}

		order, exception := o.orderService.StoreOrder(&orderRequest)

		if exception != nil {
			app.ResponseBadRequest(exception).Context(c)

			return
		}

		app.ResponseSuccess(order).Context(c)
	}
}
