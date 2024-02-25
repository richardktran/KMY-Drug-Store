package contracts

type IUserService interface {
	GetUser(id int) (interface{}, error)
}
