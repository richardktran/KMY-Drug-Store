package services

import (
	"github.com/richardktran/KMY-Drug-Store/app/models"
	repositories "github.com/richardktran/KMY-Drug-Store/app/respositories"
	"github.com/richardktran/KMY-Drug-Store/app/services/contracts"
	"github.com/richardktran/KMY-Drug-Store/pkg/app"
)

type ProductService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(
	productRepository repositories.ProductRepository,
) contracts.IProductService {
	return ProductService{
		productRepository: productRepository,
	}
}

func (s ProductService) GetProductById(id uint) (*models.Product, *app.AppError) {
	product, err := s.productRepository.GetProduct(map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s ProductService) GetProductByName(name string) (*models.Product, *app.AppError) {
	if name == "" {
		name = "<empty>"
	}

	product, err := s.productRepository.GetProduct(map[string]interface{}{"name": name})

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s ProductService) CreateProduct(data models.ProductCreation) *models.Product {
	productId, err := s.productRepository.CreateProduct(&data)

	if err != nil {
		return nil
	}

	product, err := s.GetProductById(productId)

	if err != nil {
		return nil
	}

	return product
}
