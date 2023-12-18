package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint64 `gorm:"primaryKey;unsigned"`
	Username  string `gorm:"not null;unique"`
	Email     string `gorm:"not null"`
	Name      string `gorm:"not null"`
	Deskripsi string
	Password  string `gorm:"not null"`
	NoTelp    string `gorm:"no_telp"`
	Alamat    string
	UrlGambar string `gorm:"default:'default_user.png'"`
	Restos    []Resto
	Reviews   []Review
	Koleksis  []Koleksi
}
