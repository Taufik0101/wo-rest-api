package dto

import "time"

type CreateDetailTransaksi struct {
	IdTransaksi		uint32		`json:"id_transaksi" form:"id_transaksi"`
	Customer		uint32		`json:"customer" form:"customer"`
	Seller			uint32		`json:"seller" form:"seller"`
	Produk			uint32		`json:"produk" form:"produk"`
	Pax				uint64		`json:"pax" form:"pax"`
	DownPayment		uint64		`json:"down_payment" form:"down_payment"`
	Alamat			string		`json:"alamat" form:"alamat"`
	TanggalRes		time.Time	`json:"tanggal_res" form:"tanggal_res"`
	PaymentMethods	string		`json:"payment_methods" form:"payment_methods"`
	Status			string		`json:"status" form:"status"`
}

type UpdateDetailTransaksi struct {
	IdTransaksi		uint32		`json:"id_transaksi" form:"id_transaksi"`
	Customer		uint32		`json:"customer" form:"customer"`
	Seller			uint32		`json:"seller" form:"seller"`
	Produk			uint32		`json:"produk" form:"produk"`
	Pax				uint64		`json:"pax" form:"pax"`
	DownPayment		uint64		`json:"down_payment" form:"down_payment"`
	TanggalRes		time.Time	`json:"tanggal_res" form:"tanggal_res"`
	PaymentMethods	string		`json:"payment_methods" form:"payment_methods"`
	Status			string		`json:"status" form:"status"`
}
