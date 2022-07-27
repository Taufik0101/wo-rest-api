package entity

import (
	"gorm.io/gorm"
	"time"
)

type Categories struct {
	Id_kategori		uint32		`gorm:"primary_key;auto_increment" json:"id_kategori" form:"id_kategori"`
	Nama_kategori	string		`gorm:"type:varchar(100);not null" json:"nama_kategori" form:"id_kategori"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"default:null;" json:"deleted_at" form:"deleted_at"`
}