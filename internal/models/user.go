package models

import "time"

type User struct {
	ID        uint64    `gorm:"column:id;type:uint;primaryKey"`
	Username  string    `gorm:"column:username;type:string;not null;unique"`
	Email     string    `gorm:"column:email;type:string;not null"`
	Name      string    `gorm:"column:name;type:string;not null"`
	Deskripsi string    `gorm:"column:deskripsi;type:string"`
	Password  string    `gorm:"column:password;type:string;not null"`
	NoTelp    string    `gorm:"column:no_telp;type:string"`
	Alamat    string    `gorm:"column:alamat;type:string"`
	UrlGambar string    `gorm:"column:url_gambar;type:string;default:'default_user.png'"`
	CreatedAt time.Time `gorm:"column:created_at;type:time"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:time"`
	Restos    []Resto
	Reviews   []Review
	Koleksis  []Koleksi
}
