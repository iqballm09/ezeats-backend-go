package models

import "gorm.io/gorm"

type Kecamatan struct {
	gorm.Model
	ID           uint32 `gorm:"primaryKey;unsigned"`
	KabkotaId    uint16
	Kode         uint32 `gorm:"not null;unique;unsigned"`
	Deskripsi    string `gorm:"not null"`
	RestoDetails []RestoDetail
}
