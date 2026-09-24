package model

import (
	"time"
)

type FotoProduk struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IdProduk  uint      `gorm:"column:id_produk;not null;index" json:"product_id"`
	Url       string    `gorm:"type:varchar(255);not null" json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (FotoProduk) TableName() string {
	return "foto_produk"
}
