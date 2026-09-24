package model

import (
	"time"
)

type LogProduk struct {
	ID            uint         `gorm:"primaryKey;autoIncrement" json:"id"`
	IdProduk      uint         `gorm:"column:id_produk;not null" json:"id_produk,omitempty"`
	NamaProduk    string       `gorm:"column:nama_produk;type:varchar(255);not null" json:"nama_produk"`
	Slug          string       `gorm:"type:varchar(255)" json:"slug"`
	HargaReseller int          `gorm:"column:harga_reseller;not null" json:"harga_reseler"`
	HargaKonsumen int          `gorm:"column:harga_konsumen;not null" json:"harga_konsumen"`
	Deskripsi     string       `gorm:"type:text" json:"deskripsi"`
	IdToko        uint         `gorm:"column:id_toko;not null" json:"id_toko,omitempty"`
	IdCategory    uint         `gorm:"column:id_category;not null" json:"id_category,omitempty"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`

	Toko     *Toko        `gorm:"foreignKey:IdToko" json:"toko,omitempty"`
	Category *Category    `gorm:"foreignKey:IdCategory" json:"category,omitempty"`
	Photos   []FotoProduk `gorm:"-" json:"photos,omitempty"`
}

func (LogProduk) TableName() string {
	return "log_produk"
}
