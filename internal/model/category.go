package model

import (
	"time"
)

type Category struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	NamaCategory string    `gorm:"column:nama_category;type:varchar(255);not null" json:"nama_category"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (Category) TableName() string {
	return "category"
}
