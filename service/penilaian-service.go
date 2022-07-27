package service

import (
	"github.com/Taufik0101/wo-rest-api/dto"
	"github.com/Taufik0101/wo-rest-api/entity"
	"github.com/Taufik0101/wo-rest-api/repository"
	"github.com/mashingan/smapping"
)

type PenilaianService interface {
	AllPenilaian() []entity.Penilaian
	PenilaianByVendor(id_vendor uint32) []entity.Penilaian
	SimpanPenilaian(penilaian dto.CreatePenilaian) entity.Penilaian
}

type penilaianService struct {
	penilaianRepository repository.PenilaianRepository
}

func (p penilaianService) AllPenilaian() []entity.Penilaian {
	return p.penilaianRepository.AllPenilaian()
}

func (p penilaianService) PenilaianByVendor(id_vendor uint32) []entity.Penilaian {
	return p.penilaianRepository.PenilaianByVendor(id_vendor)
}

func (p penilaianService) SimpanPenilaian(penilaian dto.CreatePenilaian) entity.Penilaian {
	newPenilaian := entity.Penilaian{}
	errSmap := smapping.FillStruct(&newPenilaian, smapping.MapFields(&penilaian))
	if errSmap != nil {
		panic("Gagal Parsing")
	}
	res := p.penilaianRepository.SimpanPenilaian(newPenilaian)
	return res
}

func NewPenilaianService(penilaianRep repository.PenilaianRepository) PenilaianService {
	return &penilaianService{
		penilaianRepository: penilaianRep,
	}
}
