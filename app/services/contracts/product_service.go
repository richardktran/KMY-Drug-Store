package contracts

import (
	"github.com/richardktran/KMY-Drug-Store/app/models"
	"github.com/richardktran/KMY-Drug-Store/pkg/app"
)

type IProductService interface {
	GetProductById(id uint) (*models.Product, *app.AppError)
	GetProductByName(name string) (*models.Product, *app.AppError)
	CreateProduct(data models.ProductCreation) *models.Product
}
