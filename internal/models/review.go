package models

type Review struct {
	ID          uint64 `gorm:"column:id;type:uint;primaryKey"`
	UserID      uint64 `gorm:"column:user_id;type:uint;foreignKey"`
	RestoID     uint64 `gorm:"column:resto_id;type:uint;foreignKey"`
	Judul       string `gorm:"not null"`
	Ulasan      string `gorm:"not null"`
	UrlGambar   string
	Rekomendasi string
	Upvotes     uint32 `gorm:"default:0;unsigned"`
	Downvotes   uint32 `gorm:"default:0;unsigned"`
}
