package repository

import (
	"github.com/Taufik0101/wo-rest-api/entity"
	"gorm.io/gorm"
)

type TransaksiRepository interface {
	AllTransaksi() []entity.Transaksi
	TransaksiByCustomer(id_customer uint32) []entity.Transaksi
	UpdateTransaksi(id_transaksi uint32, transaksi entity.Transaksi) entity.Transaksi
	CreateTransaksi(transaksi entity.Transaksi) entity.Transaksi
}

type transaksiRepository struct {
	connection *gorm.DB
}

func (t transaksiRepository) CreateTransaksi(transaksi entity.Transaksi) entity.Transaksi {
	t.connection.Save(&transaksi)
	return transaksi
}

func (t transaksiRepository) AllTransaksi() []entity.Transaksi {
	var transaksis []entity.Transaksi
	t.connection.Order("id_transaksi asc").Preload("Users").Preload("Detail").Find(&transaksis)
	return transaksis
}

func (t transaksiRepository) TransaksiByCustomer(id_customer uint32) []entity.Transaksi {
	var transaksis []entity.Transaksi
	t.connection.Where("customer = ?", id_customer).Preload("Users").Preload("Detail").Find(&transaksis)
	return transaksis
}

func (t transaksiRepository) UpdateTransaksi(id_transaksi uint32, transaksi entity.Transaksi) entity.Transaksi {
	t.connection.Where("id_transaksi = ?", id_transaksi).Updates(&transaksi)
	t.connection.Where("id_transaksi = ?", id_transaksi).Preload("Users").Preload("Detail").Find(&transaksi)
	return transaksi
}

func NewTransaksiRepository(conn *gorm.DB) TransaksiRepository {
	return &transaksiRepository{
		connection: conn,
	}
}
