package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        	uint32 	`gorm:"primary_key;auto_increment" json:"id"`
	Name  		string 	`gorm:"type:varchar(100);not null;uniqueIndex" json:"name,omitempty" form:"name"`
	Role      	string 	`gorm:"default:'User';" sql:"type:ENUM('Admin','Vendor','User')" json:"role" form:"role"`
	Email     	string 	`gorm:"type:varchar(100);not null;uniqueIndex" json:"email,omitempty" form:"email"`
	Profil		string	`gorm:"type:varchar(100)" json:"profil" form:"profil"`
	NoHP		uint64	`gorm:"default:null" json:"no_hp" form:"no_hp"`
	NIK			uint64	`gorm:"default:null" json:"nik" form:"nik"`
	Alamat		string	`gorm:"type:varchar(255)" json:"alamat" form:"alamat"`
	Toko		string	`gorm:"type:varchar(255);default:null" json:"toko" form:"toko"`		//vendor
	Category  	uint32 	`gorm:"default:null" json:"category" form:"category"`				//Vendor
	Kota      	uint32 	`gorm:"default:null" json:"kota" form:"kota"`						//vendor
	Bank1		string	`gorm:"type:varchar(100);default:null" json:"bank_1" form:"bank_1"`	//vendor
	Rekening1	uint64	`gorm:"default:null" json:"rekening_1" form:"rekening_1"`			//vendor
	Bank2		string	`gorm:"type:varchar(100);default:null" json:"bank_2" form:"bank_2"`	//vendor
	Rekening2	uint64	`gorm:"default:null" json:"rekening_2" form:"rekening_2"`			//vendor
	Rating		uint16	`gorm:"default:null" json:"rating" form:"rating"`					//vendor
	Password  	string 	`gorm:"->;<-;not null" json:"-"`
	Token     	string 	`gorm:"-" json:"token"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt `gorm:"default:null;" json:"deleted_at" form:"deleted_at"`
	Cat       	Categories     `gorm:"foreignkey:Category;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"cat"`
	Kot       	Kota           `gorm:"foreignkey:Kota;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"kot"`
}