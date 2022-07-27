package entity

import (
	"gorm.io/gorm"
	"time"
)

type Detail_Album struct {
	Id_Detail	uint32		`gorm:"primary_key;auto_increment" json:"id_detail" form:"id_detail"`
	IdAlbum		uint32		`gorm:"not null" json:"id_album" form:"id_album"`
	Foto		string		`gorm:"not null" json:"foto" form:"foto"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt 	`gorm:"default:null;" json:"deleted_at" form:"deleted_at"`
	Albums       Album         	`gorm:"foreignkey:IdAlbum;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"albums"`
}
