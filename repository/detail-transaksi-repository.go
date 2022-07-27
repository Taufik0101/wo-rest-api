package repository

import (
	"github.com/Taufik0101/wo-rest-api/entity"
	"gorm.io/gorm"
)

type DetailTransaksiRepository interface {
	DetailTransaksiSeller(id_seller uint32) []entity.Detail_Transaksi
	UpdateDetailTransaksi(id_detail uint32, detail entity.Detail_Transaksi) entity.Detail_Transaksi
	CreateDetailTransaksi(transaksi entity.Detail_Transaksi) entity.Detail_Transaksi
	DetailTransaksi(id_transaksi uint32) []entity.Detail_Transaksi
}

type detailTransaksiRepository struct {
	connection *gorm.DB
}

func (d detailTransaksiRepository) DetailTransaksi(id_transaksi uint32) []entity.Detail_Transaksi {
	var details []entity.Detail_Transaksi
	d.connection.Where("id_transaksi = ?", id_transaksi).Preload("Tx").Preload("Cus").Preload("Sell").Preload("Prod").Find(&details)
	return details
}

func (d detailTransaksiRepository) CreateDetailTransaksi(transaksi entity.Detail_Transaksi) entity.Detail_Transaksi {
	d.connection.Save(&transaksi)
	return transaksi
}

func (d detailTransaksiRepository) DetailTransaksiSeller(id_seller uint32) []entity.Detail_Transaksi {
	var details []entity.Detail_Transaksi
	d.connection.Where("seller = ?", id_seller).Preload("Tx").Preload("Cus").Preload("Sell").Preload("Prod").Find(&details)
	return details
}

func (d detailTransaksiRepository) UpdateDetailTransaksi(id_detail uint32, detail entity.Detail_Transaksi) entity.Detail_Transaksi {
	d.connection.Where("id_transaksi = ?", id_detail).Updates(&detail)
	d.connection.Where("id_transaksi = ?", id_detail).Preload("Tx").Preload("Cus").Preload("Sell").Preload("Prod").Find(&detail)
	return detail
}

func NewDetailTransaksiRepository(conn *gorm.DB) DetailTransaksiRepository {
	return &detailTransaksiRepository{
		connection: conn,
	}
}
