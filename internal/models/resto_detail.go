package models

import "gorm.io/gorm"

type RestoDetail struct {
	gorm.Model
	ID           uint64 `gorm:"primaryKey"`
	RestoID      uint64
	KecamatanID  uint32
	KabkotaID    uint16
	ProvinsiID   uint8
	Deskripsi    string `gorm:"not null"`
	Alamat       string `gorm:"not null"`
	UrlFotoMenu  string
	UrlFotoResto string  `gorm:"not null"`
	Longitude    float64 `gorm:"not null"`
	Latitude     float64 `gorm:"not null"`
	NoTelp       string
}
