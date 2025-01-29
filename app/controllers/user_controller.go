package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/richardktran/KMY-Drug-Store/app/models"
	"github.com/richardktran/KMY-Drug-Store/app/services/contracts"
	"github.com/richardktran/KMY-Drug-Store/pkg/app"
)

type UserController struct {
	userService contracts.IUserService
}

func NewUserController(
	userService contracts.IUserService,
) UserController {
	return UserController{
		userService: userService,
	}
}

func (ctl *UserController) GetUserList() func(*gin.Context) {
	return func(c *gin.Context) {
		fullName := c.DefaultQuery("full_name", "")
		phoneNumber := c.DefaultQuery("phone_number", "")

		if fullName == "" && phoneNumber == "" {
			app.ResponseBadRequest(
				app.ThrowBadRequestError(errors.New("full_name_or_phone_number_is_required"), "full_name_or_phone_number_is_required"),
			).Context(c)

			return
		}

		users, err := ctl.userService.GetUserList(fullName, phoneNumber)

		if err != nil {
			app.ResponseNotFound(
				app.ThrowNotFoundError(err, "user_not_found"),
			).Context(c)

			return
		}

		app.ResponseSuccess(users).Context(c)
	}
}

func (ctl *UserController) GetUserByPhone() func(*gin.Context) {
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
			app.ResponseNotFound(
				app.ThrowNotFoundError(err, "user_not_found"),
			).Context(c)

			return
		}

		app.ResponseSuccess(user).Context(c)

	}
}

func (ctl *UserController) AccumulateScore() func(*gin.Context) {
	return func(c *gin.Context) {
		var scoreUpdate models.ScoreUpdate

		if err := c.ShouldBind(&scoreUpdate); err != nil {
			app.ResponseBadRequest(
				app.ThrowBadRequestError(err, "invalid_request"),
			).Context(c)
			return
		}

		if scoreUpdate.PhoneNumber == "" {
			app.ResponseBadRequest(
				app.ThrowBadRequestError(errors.New("phone_number_is_required"), "phone_number_is_required"),
			).Context(c)
			return
		}

		user, err := ctl.userService.AccumulateScore(scoreUpdate.PhoneNumber)

		if err != nil {
			app.ResponseNotFound(
				app.ThrowNotFoundError(err, "user_not_found"),
			).Context(c)

			return
		}

		app.ResponseSuccess(user).Context(c)
	}
}
