package contracts

type TodoService interface {
	GetItem(id int) (interface{}, error)
}
