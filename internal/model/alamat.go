package model

import (
	"time"
)

type Alamat struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IdUser       uint      `gorm:"column:id_user;not null;index" json:"id_user,omitempty"`
	JudulAlamat  string    `gorm:"column:judul_alamat;type:varchar(255);not null" json:"judul_alamat"`
	NamaPenerima string    `gorm:"column:nama_penerima;type:varchar(255);not null" json:"nama_penerima"`
	NoTelp       string    `gorm:"column:no_telp;type:varchar(255);not null" json:"no_telp"`
	DetailAlamat string    `gorm:"column:detail_alamat;type:varchar(255);not null" json:"detail_alamat"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	User *User `gorm:"foreignKey:IdUser" json:"user,omitempty"`
}

func (Alamat) TableName() string {
	return "alamat"
}
