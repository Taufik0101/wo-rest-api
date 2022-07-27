package repository

import (
	"github.com/Taufik0101/wo-rest-api/entity"
	"gorm.io/gorm"
)

type PenilaianRepository interface {
	AllPenilaian() []entity.Penilaian
	PenilaianByVendor(id_vendor uint32) []entity.Penilaian
	SimpanPenilaian(penilaian entity.Penilaian) entity.Penilaian
}

type penilaianRepository struct {
	connection *gorm.DB
}

func (p penilaianRepository) AllPenilaian() []entity.Penilaian {
	var ratings []entity.Penilaian
	p.connection.Order("id_penilaian asc").Preload("Sell").Preload("Cus").Preload("Tx").Preload("Tx.Prod").Find(&ratings)
	return ratings
}

func (p penilaianRepository) PenilaianByVendor(id_vendor uint32) []entity.Penilaian {
	var ratings []entity.Penilaian
	p.connection.Where("seller = ?", id_vendor).Preload("Sell").Preload("Cus").Preload("Tx").Preload("Tx.Prod").Find(&ratings)
	return ratings
}

func (p penilaianRepository) SimpanPenilaian(penilaian entity.Penilaian) entity.Penilaian {
	p.connection.Save(&penilaian)
	p.connection.Preload("Sell").Preload("Cus").Preload("Tx").Preload("Tx.Prod").Find(&penilaian)
	return penilaian
}

func NewPenilaianRepository(conn *gorm.DB) PenilaianRepository {
	return &penilaianRepository{
		connection: conn,
	}
}