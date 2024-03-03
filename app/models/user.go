package models

import "github.com/richardktran/KMY-Drug-Store/pkg/model"

type User struct {
	model.BaseModel
	FullName    string `json:"full_name" gorm:"column:full_name;"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number;"`
}

type UserCreation struct {
	model.BaseModel
	FullName    string `json:"full_name" gorm:"column:full_name;"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number;"`
}

func (User) TableName() string {
	return "users"
}

func (UserCreation) TableName() string {
	return User{}.TableName()
}
