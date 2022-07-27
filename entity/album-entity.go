package entity

import (
	"gorm.io/gorm"
	"time"
)

type Album struct {
	Id_Album				uint32		`gorm:"primary_key;auto_increment" json:"id_album" form:"id_album"`
	IdVendor				uint32		`gorm:"not null" json:"id_vendor" form:"id_vendor"`
	Nama					string		`gorm:"type:varchar(255);not null" json:"nama" form:"nama"`
	Tanggal_Pelaksanaan		time.Time	`gorm:"not null" json:"tanggal_pelaksanaan" form:"tanggal_pelaksanaan"`
	CreatedAt   			time.Time
	UpdatedAt   			time.Time
	DeletedAt   			gorm.DeletedAt      `gorm:"default:null;" json:"deleted_at" form:"deleted_at"`
	Vendor					User				`gorm:"foreignkey:IdVendor;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"vendor"`
	Detail					[]Detail_Album		`gorm:"foreignKey:IdAlbum"`
}