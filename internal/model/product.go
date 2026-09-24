package model

import (
	"time"
)

type Produk struct {
	ID            uint         `gorm:"primaryKey;autoIncrement" json:"id"`
	NamaProduk    string       `gorm:"column:nama_produk;type:varchar(255);not null" json:"nama_produk"`
	Slug          string       `gorm:"type:varchar(255)" json:"slug"`
	HargaReseller int          `gorm:"column:harga_reseller;not null" json:"harga_reseler"`
	HargaKonsumen int          `gorm:"column:harga_konsumen;not null" json:"harga_konsumen"`
	Stok          int          `gorm:"not null;default:0" json:"stok"`
	Deskripsi     string       `gorm:"type:text" json:"deskripsi"`
	IdToko        uint         `gorm:"column:id_toko;not null;index" json:"id_toko,omitempty"`
	IdCategory    uint         `gorm:"column:id_category;not null;index" json:"id_category,omitempty"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`

	Toko     *Toko        `gorm:"foreignKey:IdToko" json:"toko,omitempty"`
	Category *Category    `gorm:"foreignKey:IdCategory" json:"category,omitempty"`
	Photos   []FotoProduk `gorm:"foreignKey:IdProduk" json:"photos,omitempty"`
}

func (Produk) TableName() string {
	return "produk"
}
