package services

import (
	"fmt"

	"github.com/richardktran/MyBlogBE/app/contracts"
)

type UserServiceV2 struct {
}

func NewUserServiceV2() contracts.UserService {
	return UserServiceV2{}
}

func (s UserServiceV2) GetUser(id int) (interface{}, error) {
	return fmt.Sprintf("Khoa v2 id=%d", id), nil
}
