package services

import (
	"github.com/richardktran/KMY-Drug-Store/app/models"
	repositories "github.com/richardktran/KMY-Drug-Store/app/respositories"
	"github.com/richardktran/KMY-Drug-Store/app/services/contracts"
	"github.com/richardktran/KMY-Drug-Store/pkg/app"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(
	userRepository repositories.UserRepository,
) contracts.IUserService {
	return UserService{
		userRepository: userRepository,
	}
}

func (s UserService) GetUserByPhoneNumber(phoneNumber string) (*models.User, *app.AppError) {
	user, err := s.userRepository.GetUser(map[string]interface{}{"phone_number": phoneNumber})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s UserService) GetUserById(id uint) (*models.User, *app.AppError) {
	user, err := s.userRepository.GetUser(map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s UserService) GetUserList(fullName string, phoneNumber string) ([]models.User, *app.AppError) {
	users, err := s.userRepository.GetUserList(map[string]interface{}{"full_name": fullName, "phone_number": phoneNumber})

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s UserService) CreateUser(data models.UserCreation) *models.User {
	userId, err := s.userRepository.CreateUser(&data)

	if err != nil {
		return nil
	}

	user, err := s.GetUserById(userId)

	if err != nil {
		return nil
	}

	return user
}
