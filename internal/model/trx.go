package model

import (
	"time"
)

type Trx struct {
	ID               uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	IdUser           uint        `gorm:"column:id_user;not null;index" json:"id_user,omitempty"`
	AlamatPengiriman uint        `gorm:"column:alamat_pengiriman;not null" json:"-"`
	HargaTotal       int         `gorm:"column:harga_total;not null" json:"harga_total"`
	KodeInvoice      string      `gorm:"column:kode_invoice;type:varchar(255);not null" json:"kode_invoice"`
	MethodBayar      string      `gorm:"column:method_bayar;type:varchar(255);not null" json:"method_bayar"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`

	AlamatKirim *Alamat     `gorm:"foreignKey:AlamatPengiriman" json:"alamat_kirim,omitempty"`
	DetailTrx   []DetailTrx `gorm:"foreignKey:IdTrx" json:"detail_trx,omitempty"`
}

func (Trx) TableName() string {
	return "trx"
}
