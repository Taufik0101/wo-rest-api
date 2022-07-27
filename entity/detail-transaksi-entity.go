package entity

import (
	"gorm.io/gorm"
	"time"
)

type Detail_Transaksi struct {
	Id_Detail		uint32		`gorm:"primary_key;auto_increment" json:"id_detail" form:"id_detail"`
	IdTransaksi		uint32		`gorm:"not null" json:"id_transaksi" form:"id_transaksi"`
	Customer		uint32		`gorm:"not null" json:"customer" form:"customer"`
	Seller			uint32		`gorm:"not null" json:"seller" form:"seller"`
	Produk			uint32		`gorm:"not null" json:"produk" form:"produk"`
	Pax				uint64		`gorm:"not null" json:"pax" form:"pax"`
	DownPayment		uint64		`gorm:"not null" json:"down_payment" form:"down_payment"`
	Alamat			string		`gorm:"not null" json:"alamat" form:"alamat"`
	TanggalRes		time.Time	`gorm:"not null" json:"tanggal_res" form:"tanggal_res"`
	PaymentMethods	string		`gorm:"default:null" json:"payment_methods" form:"payment_methods"`
	Status			string		`gorm:"default:null" sql:"type:ENUM('0', '1', '2')" json:"status" form:"status"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt 	`gorm:"default:null;" json:"deleted_at" form:"deleted_at"`
	Tx				Transaksi		`gorm:"foreignkey:IdTransaksi;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"tx"`
	Cus				User			`gorm:"foreignkey:Customer;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"cus"`
	Sell			User			`gorm:"foreignkey:Seller;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"sell"`
	Prod			Product			`gorm:"foreignkey:Produk;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"prod"`
}