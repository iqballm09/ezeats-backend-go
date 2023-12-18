package models

import "gorm.io/gorm"

type Koleksi struct {
	gorm.Model
	ID                   uint64 `gorm:"primaryKey;unsigned"`
	RestoID              uint64
	UserID               uint64
	IsFavoritOrBlacklist bool `gorm:"not null"`
}
