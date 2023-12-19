package models

type Provinsi struct {
	ID           uint8  `gorm:"primaryKey;unsigned"`
	Kode         uint8  `gorm:"unique;not null;unsigned"`
	Deskripsi    string `gorm:"not null"`
	Kabkotas     []Kabkota
	RestoDetails []RestoDetail
}
