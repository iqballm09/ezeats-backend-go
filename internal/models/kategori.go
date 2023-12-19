package models

type Kategori struct {
	ID        uint8  `gorm:"primaryKey;unsigned"`
	Deskripsi string `gorm:"not null"`
	Restos    []Resto
}
