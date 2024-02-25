package services

import (
	"fmt"

	"github.com/richardktran/MyBlogBE/app/services/contracts"
)

type UserServiceV2 struct {
}

func NewUserServiceV2() contracts.IUserService {
	return UserServiceV2{}
}

func (s UserServiceV2) GetUser(id int) (interface{}, error) {
	return fmt.Sprintf("Khoa v2 id=%d", id), nil
}
