package repository

import (
	"github.com/Taufik0101/wo-rest-api/entity"
	"gorm.io/gorm"
)

type ProdukRepository interface {
	AllProduk() []entity.Product
	ProdukByVendor(id_vendor uint32, role string) []entity.Product
	SimpanProduk(produk entity.Product) entity.Product
	UpdateProduk(id_produk uint32, produk entity.Product) entity.Product
	DeleteProduk(id_produk uint32)
	RandomProductLimit() []entity.Product
	ProdukById(produk_id uint32) entity.Product
}

type produkRepository struct {
	connection *gorm.DB
}

func (p produkRepository) ProdukById(produk_id uint32) entity.Product {
	var produk entity.Product
	p.connection.Where("id_produk = ?", produk_id).Preload("Users").Find(&produk)
	return produk
}

func (p produkRepository) AllProduk() []entity.Product {
	var produks []entity.Product
	p.connection.Order("id_produk asc").Preload("Users").Find(&produks)
	return produks
}

func (p produkRepository) ProdukByVendor(id_vendor uint32, role string) []entity.Product {
	var produks []entity.Product
	if role == "Vendor" {
		p.connection.Where("seller = ?", id_vendor).Preload("Users").Find(&produks)
	} else if role == "User" {
		p.connection.Where("seller = ? AND is_active = 1", id_vendor).Preload("Users").Find(&produks)

	}
	return produks
}

func (p produkRepository) SimpanProduk(produk entity.Product) entity.Product {
	p.connection.Save(&produk)
	p.connection.Preload("Users").Find(&produk)
	return produk
}

func (p produkRepository) UpdateProduk(id_produk uint32, produk entity.Product) entity.Product {
	p.connection.Where("id_produk = ?", id_produk).Updates(&produk)
	p.connection.Where("id_produk = ?", id_produk).Preload("Users").Find(&produk)
	return produk
}

func (p produkRepository) DeleteProduk(id_produk uint32) {
	p.connection.Where("id_produk = ?", id_produk).Delete(&entity.Product{})
}

func (p produkRepository) RandomProductLimit() []entity.Product {
	var produks []entity.Product
	p.connection.Raw("SELECT * FROM products WHERE is_active = 1 ORDER BY rand()").Limit(10).Preload("Users").Preload("Users.Cat").Preload("Users.Kot").Find(&produks)
	return produks
}

func NewProdukRepository(conn *gorm.DB) ProdukRepository {
	return &produkRepository{
		connection: conn,
	}
}
