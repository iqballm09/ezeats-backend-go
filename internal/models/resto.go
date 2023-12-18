package models

import "gorm.io/gorm"

type Resto struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey"`
	UserID      uint64
	KategoriID  uint8
	Nama        string `gorm:"not null"`
	HargaMin    uint64 `gorm:"not null"`
	HargaMaks   uint64 `gorm:"not null"`
	JamBuka     uint16 `gorm:"not null"`
	JamTutup    uint16 `gorm:"not null"`
	Upvotes     uint32 `gorm:"default:0"`
	Downvotes   uint32 `gorm:"default:0"`
	Reviews     []Review
	Koleksis    []Koleksi
	RestoDetail RestoDetail
}
