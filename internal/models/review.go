package models

type Review struct {
	ID          uint64 `gorm:"primaryKey;unsigned"`
	UserID      uint64
	RestoID     uint64
	Judul       string `gorm:"not null"`
	Ulasan      string `gorm:"not null"`
	UrlGambar   string
	Rekomendasi string
	Upvotes     uint32 `gorm:"default:0;unsigned"`
	Downvotes   uint32 `gorm:"default:0;unsigned"`
}
