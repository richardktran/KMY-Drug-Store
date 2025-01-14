package models

import "github.com/richardktran/KMY-Drug-Store/pkg/model"

type Order struct {
	model.BaseModel
	UserID    uint     `json:"-" gorm:"column:user_id;"`
	User      *User    `json:"user,omitempty" gorm:"foreignKey:UserID;"` // foreign key
	ProductID uint     `json:"-" gorm:"column:product_id;"`
	Product   *Product `json:"product,omitempty" gorm:"foreignKey:ProductID;"` // foreign key
	Amount    int      `json:"amount" gorm:"column:amount;"`
	Quantity  int      `json:"quantity" gorm:"column:quantity;"`
	Note      string   `json:"note,omitempty" gorm:"column:note;"`
}

type OrderMetaData struct {
	Total int64 `json:"total,omitempty"`
}

type OrderCreation struct {
	model.BaseModel
	PhoneNumber string `json:"phone_number" gorm:"-"`
	FullName    string `json:"full_name,omitempty" gorm:"-"`
	UserId      uint   `json:"-" gorm:"column:user_id;"`
	ProductName string `json:"product_name" gorm:"-"`
	Unit        string `json:"unit,omitempty" gorm:"-"`
	ProductId   uint   `json:"-" gorm:"column:product_id;"`
	Quantity    int    `json:"quantity,omitempty" gorm:"column:quantity;"`
	Amount      int    `json:"amount" gorm:"column:amount;"`
	Score       int    `json:"score,omitempty" gorm:"-"`
	Note        string `json:"note,omitempty" gorm:"column:note;default:null"`
}

func (Order) TableName() string {
	return "orders"
}

func (OrderCreation) TableName() string {
	return Order{}.TableName()
}

func (OrderMetaData) TableName() string {
	return Order{}.TableName()
}
