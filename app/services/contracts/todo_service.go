package contracts

type ITodoService interface {
	GetItem(id int) (interface{}, error)
}
