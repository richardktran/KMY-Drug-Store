package repositories

import (
	model "github.com/richardktran/MyBlogBE/app/models"
	"github.com/richardktran/MyBlogBE/pkg/app"
	"github.com/richardktran/MyBlogBE/pkg/database"
	"gorm.io/gorm"
)

type TodoRepository struct {
}

func NewTodoRepository() TodoRepository {
	return TodoRepository{}
}

func (r TodoRepository) GetItem(condition map[string]interface{}) (*model.TodoItem, *app.AppError) {
	var data model.TodoItem
	db := database.GetDB()

	if err := db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, app.ThrowDefaultNotFoundError(err)
		} else {
			return nil, app.ThrowInternalServerError(err)
		}
	}

	return &data, nil
}
