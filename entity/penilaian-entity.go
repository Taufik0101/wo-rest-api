package entity

import (
	"gorm.io/gorm"
	"time"
)

type Penilaian struct {
	Id_Penilaian		uint32		`gorm:"primary_key;auto_increment" json:"id_penilaian" form:"id_penilaian"`
	Seller				uint32		`gorm:"not null" json:"seller" form:"seller"`
	Customer			uint32		`gorm:"not null" json:"customer" form:"customer"`
	Trans				uint32		`gorm:"not null" json:"trans" form:"trans"`
	Judul				string		`gorm:"type:varchar(255);not null" json:"judul" form:"judul"`
	Pesan				string		`gorm:"type:longtext" json:"pesan" form:"pesan"`
	Star				uint16		`gorm:"not null" json:"star" form:"star"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt 		`gorm:"default:null;" json:"deleted_at" form:"deleted_at"`
	Sell		User				`gorm:"foreignkey:Seller;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"sell"`
	Cus			User				`gorm:"foreignkey:Customer;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"cus"`
	Tx			Detail_Transaksi	`gorm:"foreignkey:Trans;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"tx"`
}