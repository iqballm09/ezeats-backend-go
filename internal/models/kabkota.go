package models

type Kabkota struct {
	ID           uint16 `gorm:"primaryKey;unsigned"`
	ProvinsiID   uint8
	Kode         uint16 `gorm:"unique;not null;unsigned"`
	Deskripsi    string `gorm:"not null"`
	Kecamatans   []Kecamatan
	RestoDetails []RestoDetail
}
