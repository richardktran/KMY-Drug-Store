package repositories

import (
	"github.com/richardktran/KMY-Drug-Store/app/models"
	"github.com/richardktran/KMY-Drug-Store/pkg/app"
	"github.com/richardktran/KMY-Drug-Store/pkg/database"
)

type OrderRepository struct {
}

func NewOrderRepository() OrderRepository {
	return OrderRepository{}
}

func (r OrderRepository) GetOrder(condition map[string]interface{}) (*models.Order, *app.AppError) {
	db := database.GetDB()
	var order models.Order

	if err := db.Preload("User").Preload("Product").Where(condition).First(&order).Error; err != nil {
		return nil, app.ThrowDefaultNotFoundError(err)
	}

	return &order, nil
}

func (r OrderRepository) StoreOrder(data *models.OrderCreation) (uint, *app.AppError) {
	db := database.GetDB()

	if err := db.Create(&data).Error; err != nil {
		return 0, app.ThrowInternalServerError(err)
	}
	return data.ID, nil
}
