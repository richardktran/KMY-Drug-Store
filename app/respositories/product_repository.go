package repositories

import (
	"github.com/richardktran/KMY-Drug-Store/app/models"
	"github.com/richardktran/KMY-Drug-Store/pkg/app"
	"github.com/richardktran/KMY-Drug-Store/pkg/database"
)

type ProductRepository struct {
}

func NewProductRepository() ProductRepository {
	return ProductRepository{}
}

func (r ProductRepository) GetProduct(condition map[string]any) (*models.Product, *app.AppError) {
	db := database.GetDB()
	var product models.Product

	if err := db.Where(condition).First(&product).Error; err != nil {
		return nil, app.ThrowDefaultNotFoundError(err)
	}

	return &product, nil
}

func (r ProductRepository) CreateProduct(data *models.ProductCreation) (uint, *app.AppError) {
	db := database.GetDB()

	// Create and get user id
	if err := db.Create(&data).Error; err != nil {
		return 0, app.ThrowInternalServerError(err)
	}

	return data.ID, nil
}
