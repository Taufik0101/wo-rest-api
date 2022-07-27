package entity

import (
	"gorm.io/gorm"
	"time"
)

type Request struct {
	ReqId       uint32 	`gorm:"primary_key;auto_increment" json:"id_req" form:"id_req"`
	Name  		string 	`gorm:"size:100;not null;unique_index" json:"nama,omitempty" form:"nama"`
	Role      	string 	`gorm:"default:'Vendor';" sql:"type:ENUM('Admin', 'User', 'Vendor')" json:"role" form:"role"`
	Email     	string 	`gorm:"size:100;not null;unique_index" json:"email,omitempty" form:"email"`
	NoHP		uint64	`gorm:"default:null" json:"no_hp" form:"no_hp"`
	NIK			uint64	`gorm:"default:null" json:"nik" form:"nik"`
	Category  	uint32 	`gorm:"not null;" json:"id_kategori" form:"id_kategori"`
	KTP			string 	`gorm:"type:varchar(255);not null" json:"ktp" form:"ktp"`
	Selfie    	string 	`gorm:"type:varchar(255);not null" json:"selfie_ktp" form:"selfie_ktp"`
	Status      string 	`gorm:"default: null;" sql:"type:ENUM('0', '1')" json:"status" form:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"default:null;" json:"deleted_at" form:"deleted_at"`
	Cat       Categories     `gorm:"foreignkey:Category;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"cat"`
}