package entity

import (
	"gorm.io/gorm"
	"time"
)

type Transaksi struct {
	Id_Transaksi	uint32		`gorm:"primary_key;auto_increment" json:"id_transaksi" form:"id_transaksi"`
	Customer		uint32		`gorm:"not null" json:"customer" form:"customer"`
	TotalPrice		uint64		`gorm:"not null" json:"total_price" form:"total_price"`
	PaymentMethods	string		`gorm:"default:null" json:"payment_methods" form:"payment_methods"`
	Status			string		`gorm:"default:null" sql:"type:ENUM('0', '1', '2')" json:"status" form:"status"`
	IsReview		string		`gorm:"default:null" json:"is_review" form:"is_review"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt 		`gorm:"default:null;" json:"deleted_at" form:"deleted_at"`
	Users			User				`gorm:"foreignkey:Customer;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"users"`
	Detail			*[]Detail_Transaksi	`gorm:"-" json:"detail,omitempty"`
}