package contracts

import (
	"github.com/richardktran/KMY-Drug-Store/app/models"
	"github.com/richardktran/KMY-Drug-Store/pkg/app"
)

type IUserService interface {
	GetUserByPhoneNumber(phoneNumber string) (*models.User, *app.AppError)
	GetUserList(fullName string, phoneNumber string) ([]models.User, *app.AppError)
	CreateUser(data models.UserCreation) *models.User
	UpdateUserById(id uint, data models.UserUpdate) *models.User
}
