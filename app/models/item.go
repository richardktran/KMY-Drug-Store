package model

import "github.com/richardktran/MyBlogBE/pkg/model"

type TodoItem struct {
	model.BaseModel
	Title       string     `json:"title" gorm:"column:title;"`
	Description string     `json:"description" gorm:"column:description;"`
	Status      ItemStatus `json:"status" gorm:"column:status;"`
}

type TodoItemCreation struct {
	Id          int    `json:"-" gorm:"column:id;"`
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
}

type TodoItemUpdate struct {
	Title       *string     `json:"title" gorm:"column:title;"` // pointer to allow null or empty string
	Description *string     `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
}

func (TodoItem) TableName() string {
	return "todo_items"
}

func (TodoItemCreation) TableName() string {
	return TodoItem{}.TableName()
}

func (TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}
