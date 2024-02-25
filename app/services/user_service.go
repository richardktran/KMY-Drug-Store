package services

import (
	"fmt"

	"github.com/richardktran/MyBlogBE/app/contracts"
)

type UserService struct {
}

func NewUserService() contracts.UserService {
	return UserService{}
}

func (s UserService) GetUser(id int) (interface{}, error) {
	return fmt.Sprintf("Khoa id=%d", id), nil
}
