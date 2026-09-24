package model

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama         string    `gorm:"type:varchar(255);not null" json:"nama"`
	KataSandi    string    `gorm:"column:kata_sandi;type:varchar(255);not null" json:"-"`
	NoTelp       string    `gorm:"column:notelp;type:varchar(255);uniqueIndex;not null" json:"no_telp"`
	TanggalLahir string    `gorm:"column:tanggal_lahir;type:varchar(255)" json:"tanggal_lahir"`
	JenisKelamin string    `gorm:"column:jenis_kelamin;type:varchar(255)" json:"jenis_kelamin"`
	Tentang      string    `gorm:"type:text" json:"tentang"`
	Pekerjaan    string    `gorm:"type:varchar(255)" json:"pekerjaan"`
	Email        string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	IdProvinsi   string    `gorm:"column:id_provinsi;type:varchar(255)" json:"id_provinsi"`
	IdKota       string    `gorm:"column:id_kota;type:varchar(255)" json:"id_kota"`
	IsAdmin      bool      `gorm:"column:is_admin;default:false" json:"is_admin"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Toko    *Toko    `gorm:"foreignKey:IdUser" json:"toko,omitempty"`
	Alamats []Alamat `gorm:"foreignKey:IdUser" json:"alamats,omitempty"`
}

func (User) TableName() string {
	return "users"
}
