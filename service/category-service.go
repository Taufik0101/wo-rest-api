package service

import (
	"github.com/Taufik0101/wo-rest-api/dto"
	"github.com/Taufik0101/wo-rest-api/entity"
	"github.com/Taufik0101/wo-rest-api/repository"
	"github.com/mashingan/smapping"
)

type CategoriService interface {
	AllKategori() []entity.Categories
	SimpanKategori(kategori dto.CreateKategori) entity.Categories
	UpdateKategori(id_kategori uint32, kategori dto.UpdateKategori) entity.Categories
	HapusKategori(id_kategori uint32)
}

type categoryService struct {
	categoryRepository repository.KategoriRepository
}

func (c categoryService) AllKategori() []entity.Categories {
	return c.categoryRepository.AllKategori()
}

func (c categoryService) SimpanKategori(kategori dto.CreateKategori) entity.Categories {
	newKategori := entity.Categories{}
	errSmap := smapping.FillStruct(&newKategori, smapping.MapFields(&kategori))
	if errSmap != nil {
		panic("Gagal Parsing")
	}
	res := c.categoryRepository.SimpanKategori(newKategori)
	return res
}

func (c categoryService) UpdateKategori(id_kategori uint32, kategori dto.UpdateKategori) entity.Categories {
	upKategori := entity.Categories{}
	errSmap := smapping.FillStruct(&upKategori, smapping.MapFields(&kategori))
	if errSmap != nil {
		panic("Gagal Parsing")
	}else {
		res := c.categoryRepository.UpdateKategori(id_kategori, upKategori)
		return res
	}
}

func (c categoryService) HapusKategori(id_kategori uint32) {
	c.categoryRepository.HapusKategori(id_kategori)
}

func NewKategoriService(categoryRep repository.KategoriRepository) CategoriService {
	return &categoryService{
		categoryRepository: categoryRep,
	}
}
