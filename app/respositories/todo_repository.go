package repositories

import (
	model "github.com/richardktran/KMY-Drug-Store/app/models"
	"github.com/richardktran/KMY-Drug-Store/pkg/app"
	"github.com/richardktran/KMY-Drug-Store/pkg/database"
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
