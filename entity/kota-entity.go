package entity

import (
	"time"
)

type Kota struct {
	Id_Kota		uint32	`gorm:"primary_key;auto_increment" json:"id_kota" form:"id_kota"`
	Nama_Kota	string	`gorm:"type:varchar(100);not null;uniqueIndex" json:"nama_kota" form:"nama_kota"`
	CreatedAt time.Time
	UpdatedAt time.Time
	//DeletedAt gorm.DeletedAt `gorm:"default:null;" json:"deleted_at" form:"deleted_at"`
}