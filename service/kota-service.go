package service

import (
	"github.com/Taufik0101/wo-rest-api/dto"
	"github.com/Taufik0101/wo-rest-api/entity"
	"github.com/Taufik0101/wo-rest-api/repository"
	"github.com/mashingan/smapping"
	"strings"
)

type KotaService interface {
	AllKota() []entity.Kota
	SimpanKota(kota dto.CreateKota) entity.Kota
	UpdateKota(id_kota uint32, kota dto.UpdateKota) entity.Kota
	DeleteKota(id_kota uint32)
	FindKotaByName(name string) bool
}

type kotaService struct {
	kotaRepository repository.KotaRepository
}

func (k kotaService) FindKotaByName(name string) bool {
	res := k.kotaRepository.FindByName(name)
	return !(res.Error == nil)
}

func (k kotaService) AllKota() []entity.Kota {
	return k.kotaRepository.AllKota()
}

func (k kotaService) SimpanKota(kota dto.CreateKota) entity.Kota {
	newKota := entity.Kota{}
	errSmap := smapping.FillStruct(&newKota, smapping.MapFields(&kota))
	if errSmap != nil {
		panic("Gagal Parsing")
	}
	newKota.Nama_Kota = strings.Title(newKota.Nama_Kota)
	res := k.kotaRepository.SimpanKota(newKota)
	return res
}

func (k kotaService) UpdateKota(id_kota uint32, kota dto.UpdateKota) entity.Kota {
	upKota := entity.Kota{}
	errSmap := smapping.FillStruct(&upKota, smapping.MapFields(&kota))
	if errSmap != nil {
		panic("Gagal Parsing")
	}else {
		res := k.kotaRepository.UpdateKota(id_kota, upKota)
		return res
	}
}

func (k kotaService) DeleteKota(id_kota uint32) {
	k.kotaRepository.DeleteKota(id_kota)
}

func NewKotaService(kotaRep repository.KotaRepository) KotaService {
	return &kotaService{
		kotaRepository: kotaRep,
	}
}
