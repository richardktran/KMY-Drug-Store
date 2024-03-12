package models

import "github.com/richardktran/KMY-Drug-Store/pkg/model"

type User struct {
	model.BaseModel
	FullName    string `json:"full_name" gorm:"column:full_name;"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number;"`
	ScoreUsed   int    `json:"score_used" gorm:"column:score_used;"`
	RemainScore int    `json:"remain_score" gorm:"-"`
	MaxScore    int    `json:"max_score" gorm:"-"`
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
