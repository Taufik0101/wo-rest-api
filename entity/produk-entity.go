package entity

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	Id_produk		uint32		`gorm:"primary_key;auto_increment" json:"id_produk" form:"id_produk"`
	Seller			uint32		`gorm:"not null" json:"seller" form:"seller"`
	Nama_Produk		string		`gorm:"type:varchar(200);not null" json:"nama_produk" form:"nama_produk"`
	Harga			string		`gorm:"not null" json:"harga" form:"harga"`
	DP				uint16		`gorm:"not null" json:"dp" form:"dp"`
	Image			string		`gorm:"type:varchar(150);not null" json:"image" form:"image"`
	Deskripsi		string		`gorm:"type:longtext;not null" json:"deskripsi" form:"deskripsi"`
	Is_Active		string		`gorm:"default:'1'" sql:"type:ENUM('0', '1')" json:"is_active" form:"is_active"`
	CreatedAt   	time.Time
	UpdatedAt   	time.Time
	DeletedAt   	gorm.DeletedAt      `gorm:"default:null;" json:"deleted_at" form:"deleted_at"`
	Users			User				`gorm:"foreignkey:Seller;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"users"`
}