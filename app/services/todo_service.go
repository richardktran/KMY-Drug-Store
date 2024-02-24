package services

import (
	repositories "github.com/richardktran/MyBlogBE/app/respositories"
	"github.com/richardktran/MyBlogBE/pkg/app"
)

type TodoService struct {
	repository repositories.TodoRepository
}

func NewTodoService(repository repositories.TodoRepository) TodoService {
	return TodoService{
		repository: repository,
	}
}

func (s TodoService) GetItem(id int) (interface{}, error) {
	data, err := s.repository.GetItem(map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return nil, app.ThrowNotFoundError(err, "item_not_found")
	}

	return data, nil
}
