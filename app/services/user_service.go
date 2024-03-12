package services

import (
	"github.com/richardktran/KMY-Drug-Store/app/models"
	repositories "github.com/richardktran/KMY-Drug-Store/app/respositories"
	"github.com/richardktran/KMY-Drug-Store/app/services/contracts"
	"github.com/richardktran/KMY-Drug-Store/conf"
	"github.com/richardktran/KMY-Drug-Store/pkg/app"
)

type UserService struct {
	userRepository  repositories.UserRepository
	orderRepository repositories.OrderRepository
}

func NewUserService(
	userRepository repositories.UserRepository,
	orderRepository repositories.OrderRepository,
) contracts.IUserService {
	return UserService{
		userRepository:  userRepository,
		orderRepository: orderRepository,
	}
}

func (s UserService) GetUserByPhoneNumber(phoneNumber string) (*models.User, *app.AppError) {
	user, err := s.userRepository.GetUser(map[string]interface{}{"phone_number": phoneNumber})

	if err != nil {
		return nil, err
	}

	user.RemainScore = s.CalculateUserScore(user)
	user.MaxScore = s.CalculateMaximumScoreUsed(user)

	return user, nil
}

func (s UserService) GetUserById(id uint) (*models.User, *app.AppError) {
	user, err := s.userRepository.GetUser(map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	user.RemainScore = s.CalculateUserScore(user)
	user.MaxScore = s.CalculateMaximumScoreUsed(user)

	return user, nil
}

func (s UserService) GetUserList(fullName string, phoneNumber string) ([]models.User, *app.AppError) {
	users, err := s.userRepository.GetUserList(map[string]interface{}{"full_name": fullName, "phone_number": phoneNumber})

	if err != nil {
		return nil, err
	}

	for i := range users {
		users[i].RemainScore = s.CalculateUserScore(&users[i])
		users[i].MaxScore = s.CalculateMaximumScoreUsed(&users[i])
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

func (s UserService) CalculateUserScore(user *models.User) int {
	totalAmount, error := s.orderRepository.GetTotalAmountOfUser(user.ID)

	if error != nil {
		return 0
	}

	return int(totalAmount/1000) - user.ScoreUsed*conf.SCORE_RATE
}

func (s UserService) CalculateMaximumScoreUsed(user *models.User) int {
	return s.CalculateUserScore(user) / conf.SCORE_RATE
}
