package services

import (
	repositories "github.com/richardktran/MyBlogBE/app/respositories"
	"github.com/richardktran/MyBlogBE/app/services/contracts"
	"github.com/richardktran/MyBlogBE/pkg/app"
)

type TodoService struct {
	repository repositories.TodoRepository
}

func NewTodoService(repository repositories.TodoRepository) contracts.ITodoService {
	return TodoService{
		repository: repository,
	}
}

func (s TodoService) GetItem(id int) (interface{}, *app.AppError) {
	data, err := s.repository.GetItem(map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return nil, err
	}

	return data, nil
}
