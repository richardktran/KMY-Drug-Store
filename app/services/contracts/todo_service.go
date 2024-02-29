package contracts

import "github.com/richardktran/KMY-Drug-Store/pkg/app"

type ITodoService interface {
	GetItem(id int) (interface{}, *app.AppError)
}
