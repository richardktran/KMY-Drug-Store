package models

import "github.com/richardktran/KMY-Drug-Store/pkg/model"

type Product struct {
	model.BaseModel
	Name string `json:"name" gorm:"column:name;"`
	Unit string `json:"unit" gorm:"column:unit;"`
}

type ProductCreation struct {
	model.BaseModel
	Name string `json:"name" gorm:"column:name;"`
	Unit string `json:"unit" gorm:"column:unit;default:null"`
}

func (Product) TableName() string {
	return "products"
}

func (ProductCreation) TableName() string {
	return Product{}.TableName()
}
