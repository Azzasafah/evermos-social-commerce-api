package model

import (
	"time"
)

type DetailTrx struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	IdTrx       uint       `gorm:"column:id_trx;not null;index" json:"id_trx,omitempty"`
	IdLogProduk uint       `gorm:"column:id_log_produk;not null;index" json:"id_log_produk,omitempty"`
	IdToko      uint       `gorm:"column:id_toko;not null;index" json:"id_toko,omitempty"`
	Kuantitas   int        `gorm:"not null" json:"kuantitas"`
	HargaTotal  int        `gorm:"column:harga_total;not null" json:"harga_total"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`

	Product *LogProduk `gorm:"foreignKey:IdLogProduk" json:"product,omitempty"`
	Toko    *Toko      `gorm:"foreignKey:IdToko" json:"toko,omitempty"`
}

func (DetailTrx) TableName() string {
	return "detail_trx"
}
