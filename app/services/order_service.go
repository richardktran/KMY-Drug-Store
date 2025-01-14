package services

import (
	"github.com/richardktran/KMY-Drug-Store/app/models"
	repositories "github.com/richardktran/KMY-Drug-Store/app/respositories"
	"github.com/richardktran/KMY-Drug-Store/app/services/contracts"
	"github.com/richardktran/KMY-Drug-Store/pkg/app"
)

type OrderService struct {
	userService    contracts.IUserService
	productService contracts.IProductService
	repository     repositories.OrderRepository
}

func NewOrderService(
	repository repositories.OrderRepository,
	userService contracts.IUserService,
	productService contracts.IProductService,
) contracts.IOrderService {
	return OrderService{
		repository:     repository,
		userService:    userService,
		productService: productService,
	}
}

func (s OrderService) GetOrderById(id uint) (*models.Order, *app.AppError) {
	order, err := s.repository.GetOrder(map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s OrderService) GetAllOrders(condition map[string]interface{}, recursive bool) ([]models.Order, *app.Meta, *app.AppError) {
	orders, metaData, err := s.repository.GetAllOrders(condition, recursive)

	if err != nil {
		return nil, nil, err
	}

	return orders, metaData, nil
}

func (s OrderService) StoreOrder(data *models.OrderCreation) (*models.Order, *app.AppError) {
	user, err := s.userService.GetUserByPhoneNumber(data.PhoneNumber)

	if err != nil {
		userDataCreation := models.UserCreation{
			PhoneNumber: data.PhoneNumber,
			FullName:    data.FullName,
		}
		user = s.userService.CreateUser(userDataCreation)
	}

	if data.Score > 0 {
		if user.MaxScore < data.Score {
			return nil, app.ThrowBadRequestError(nil, "not_enough_score")
		}

		userDataUpdate := models.UserUpdate{
			ScoreUsed: user.ScoreUsed + data.Score,
		}

		userUpdated := s.userService.UpdateUserById(user.ID, userDataUpdate)

		if userUpdated == nil {
			return nil, app.ThrowBadRequestError(
				nil,
				"update_score_failed",
			)
		}

	}

	if user == nil {
		return nil, app.ThrowBadRequestError(nil, "user_not_found")
	}

	// Check if product exists
	product, err := s.productService.GetProductByName(data.ProductName)

	if err != nil {
		productDataCreation := models.ProductCreation{
			Name: data.ProductName,
			Unit: data.Unit,
		}
		product = s.productService.CreateProduct(productDataCreation)
	}

	data.UserId = user.ID
	data.ProductId = product.ID

	orderId, err := s.repository.StoreOrder(data)

	if err != nil {
		return nil, app.ThrowBadRequestError(err, "store_order_failed")
	}

	order, err := s.GetOrderById(orderId)

	if err != nil {
		return nil, app.ThrowBadRequestError(err, "store_order_failed")
	}

	return order, nil
}
