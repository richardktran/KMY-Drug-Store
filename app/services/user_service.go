package services

import "fmt"

type UserService struct {
}

func NewUserService() UserService {
	return UserService{}
}

func (s UserService) GetUser(id int) (interface{}, error) {
	return fmt.Sprintf("Khoa id=%d", id), nil
}
