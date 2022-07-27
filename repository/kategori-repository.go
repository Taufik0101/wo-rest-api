package repository

import (
	"github.com/Taufik0101/wo-rest-api/entity"
	"gorm.io/gorm"
)

type KategoriRepository interface {
	AllKategori() []entity.Categories
	SimpanKategori(kategori entity.Categories) entity.Categories
	UpdateKategori(id_kategori uint32, kategori entity.Categories) entity.Categories
	HapusKategori(id_kategori uint32)
}

type kategoriRepository struct {
	connection *gorm.DB
}

func (k kategoriRepository) AllKategori() []entity.Categories {
	var categories []entity.Categories
	k.connection.Order("id_kategori asc").Find(&categories)
	return categories
}

func (k kategoriRepository) SimpanKategori(kategori entity.Categories) entity.Categories {
	k.connection.Save(&kategori)
	return kategori
}

func (k kategoriRepository) UpdateKategori(id_kategori uint32, kategori entity.Categories) entity.Categories {
	k.connection.Where("id_kategori = ?", id_kategori).Updates(&kategori)
	return kategori
}

func (k kategoriRepository) HapusKategori(id_kategori uint32) {
	k.connection.Where("id_kategori = ?", id_kategori).Delete(&entity.Categories{})
}

func NewKategoriRepository(conn *gorm.DB) KategoriRepository {
	return &kategoriRepository{
		connection: conn,
	}
}
