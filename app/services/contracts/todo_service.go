package contracts

import "github.com/richardktran/MyBlogBE/pkg/app"

type ITodoService interface {
	GetItem(id int) (interface{}, *app.AppError)
}
