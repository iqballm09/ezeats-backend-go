package models

type Koleksi struct {
	ID                   uint64 `gorm:"primaryKey;unsigned"`
	RestoID              uint64
	UserID               uint64
	IsFavoritOrBlacklist bool `gorm:"not null"`
}
