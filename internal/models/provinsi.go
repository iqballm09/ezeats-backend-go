package models

import "gorm.io/gorm"

type Provinsi struct {
	gorm.Model
	ID           uint8  `gorm:"primaryKey;unsigned"`
	Kode         uint8  `gorm:"unique;not null;unsigned"`
	Deskripsi    string `gorm:"not null"`
	Kabkotas     []Kabkota
	RestoDetails []RestoDetail
}
