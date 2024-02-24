package model

type BaseModel struct {
	ID        uint   `json:"id" gorm:"primary_key;column:id"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"`
}
