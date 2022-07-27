package service

import (
	"github.com/Taufik0101/wo-rest-api/dto"
	"github.com/Taufik0101/wo-rest-api/entity"
	"github.com/Taufik0101/wo-rest-api/repository"
	"github.com/mashingan/smapping"
)

type ProdukService interface {
	AllProduk() []entity.Product
	ProdukByVendor(id_vendor uint32, role string) []entity.Product
	SimpanProduk(produk dto.CreateProduk) entity.Product
	UpdateProduk(id_produk uint32, produk dto.UpdateProduk) entity.Product
	DeleteProduk(id_produk uint32)
	RandomProductLimit() []entity.Product
	ProdukById(id_produk uint32) entity.Product
}

type produkService struct {
	produkRepository repository.ProdukRepository
}

func (p produkService) ProdukById(id_produk uint32) entity.Product {
	return p.produkRepository.ProdukById(id_produk)
}

func (p produkService) AllProduk() []entity.Product {
	return p.produkRepository.AllProduk()
}

func (p produkService) ProdukByVendor(id_vendor uint32, role string) []entity.Product {
	return p.produkRepository.ProdukByVendor(id_vendor, role)
}

func (p produkService) SimpanProduk(produk dto.CreateProduk) entity.Product {
	newProduk := entity.Product{}
	errSmap := smapping.FillStruct(&newProduk, smapping.MapFields(&produk))
	if errSmap != nil {
		panic("Gagal Parsing")
	}
	res := p.produkRepository.SimpanProduk(newProduk)
	return res
}

func (p produkService) UpdateProduk(id_produk uint32, produk dto.UpdateProduk) entity.Product {
	upProduk := entity.Product{}
	errSmap := smapping.FillStruct(&upProduk, smapping.MapFields(&produk))
	if errSmap != nil {
		panic("Gagal Parsing")
	}else {
		res := p.produkRepository.UpdateProduk(id_produk, upProduk)
		return res
	}
}

func (p produkService) DeleteProduk(id_produk uint32) {
	p.produkRepository.DeleteProduk(id_produk)
}

func (p produkService) RandomProductLimit() []entity.Product {
	return p.produkRepository.RandomProductLimit()
}

func NewProdukService(produkRep repository.ProdukRepository) ProdukService {
	return &produkService{
		produkRepository: produkRep,
	}
}
