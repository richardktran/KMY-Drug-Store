package contracts

type UserService interface {
	GetUser(id int) (interface{}, error)
}
