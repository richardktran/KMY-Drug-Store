package contracts

import (
	"github.com/richardktran/KMY-Drug-Store/app/models"
	"github.com/richardktran/KMY-Drug-Store/pkg/app"
)

type IOrderService interface {
	GetOrderById(id uint) (*models.Order, *app.AppError)
	GetAllOrders(condition map[string]interface{}, recursive bool) ([]models.Order, *app.AppError)
	StoreOrder(data *models.OrderCreation) (*models.Order, *app.AppError)
}
