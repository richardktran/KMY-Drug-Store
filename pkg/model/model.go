package model

import (
	"database/sql/driver"
	"time"
)

type NullTime struct {
	Time  time.Time
	Valid bool `json:"-" gorm:"-"`
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

type BaseModel struct {
	ID        uint     `json:"id" gorm:"primary_key;column:id"`
	CreatedAt NullTime `json:"created_at" gorm:"column:created_at"`
	UpdatedAt NullTime `json:"updated_at" gorm:"column:updated_at"`
}
