package services

import (
	"fmt"

	"github.com/richardktran/MyBlogBE/app/services/contracts"
)

type UserService struct {
}

func NewUserService() contracts.IUserService {
	return UserService{}
}

func (s UserService) GetUser(id int) (interface{}, error) {
	return fmt.Sprintf("Khoa id=%d", id), nil
}
