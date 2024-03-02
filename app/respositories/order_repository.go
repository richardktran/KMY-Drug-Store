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

func (r OrderRepository) GetAllOrders(condition map[string]interface{}, recursive bool) ([]models.Order, *app.Meta, *app.AppError) {
	var orders []models.Order
	var metaData models.OrderMetaData
	db := database.GetDB().Where(condition)

	if recursive {
		db = db.Joins("User").Joins("Product")
	} else {
		// set user and product to nil
		for i := range orders {
			orders[i].User = nil
			orders[i].Product = nil
		}
	}

	if err := db.Order("orders.created_at DESC").
		Find(&orders).Error; err != nil {
		return nil, nil, app.ThrowDefaultNotFoundError(err)
	}

	// Query total amount
	if err := db.Select("SUM(orders.amount)").Scan(&metaData.Total).Error; err != nil {
		return nil, nil, app.ThrowDefaultNotFoundError(err)
	}

	return orders, &app.Meta{Data: metaData}, nil
}

func (r OrderRepository) StoreOrder(data *models.OrderCreation) (uint, *app.AppError) {
	db := database.GetDB()

	if err := db.Create(&data).Error; err != nil {
		return 0, app.ThrowInternalServerError(err)
	}
	return data.ID, nil
}
