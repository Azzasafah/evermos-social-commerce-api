package model

import (
	"time"
)

type Toko struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IdUser    uint      `gorm:"column:id_user;not null;index" json:"user_id"`
	NamaToko  string    `gorm:"column:nama_toko;type:varchar(255);not null" json:"nama_toko"`
	UrlFoto   string    `gorm:"column:url_foto;type:varchar(255)" json:"url_foto"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User *User `gorm:"foreignKey:IdUser" json:"user,omitempty"`
}

func (Toko) TableName() string {
	return "toko"
}
