package models

import "gorm.io/gorm"

type Kategori struct {
	gorm.Model
	ID        uint8  `gorm:"primaryKey;unsigned"`
	Deskripsi string `gorm:"not null"`
	Restos    []Resto
}
