package service

import (
	"github.com/Taufik0101/wo-rest-api/dto"
	"github.com/Taufik0101/wo-rest-api/entity"
	"github.com/Taufik0101/wo-rest-api/repository"
	"github.com/mashingan/smapping"
)

type TransaksiService interface {
	AllTransaksi() []entity.Transaksi
	TransaksiByCustomer(id_customer uint32) []entity.Transaksi
	UpdateTransaksi(id_transaksi uint32, transaksi dto.UpdateTransaksi) entity.Transaksi
	CreateTransaksi(transaksi entity.Transaksi) entity.Transaksi
}

type transaksiService struct {
	transaksiRepository repository.TransaksiRepository
}

func (t transaksiService) CreateTransaksi(transaksi entity.Transaksi) entity.Transaksi {
	newTransaksi := entity.Transaksi{}
	errSmap := smapping.FillStruct(&newTransaksi, smapping.MapFields(&transaksi))
	if errSmap != nil {
		panic("Gagal Parsing")
	}
	res := t.transaksiRepository.CreateTransaksi(newTransaksi)
	return res
}

func (t transaksiService) AllTransaksi() []entity.Transaksi {
	return t.transaksiRepository.AllTransaksi()
}

func (t transaksiService) TransaksiByCustomer(id_customer uint32) []entity.Transaksi {
	return t.transaksiRepository.TransaksiByCustomer(id_customer)
}

func (t transaksiService) UpdateTransaksi(id_transaksi uint32, transaksi dto.UpdateTransaksi) entity.Transaksi {
	upTransaksi := entity.Transaksi{}
	errSmap := smapping.FillStruct(&upTransaksi, smapping.MapFields(&transaksi))
	if errSmap != nil {
		panic("Gagal Parsing")
	}else {
		res := t.transaksiRepository.UpdateTransaksi(id_transaksi, upTransaksi)
		return res
	}
}

func NewTransaksiService(transaksiRep repository.TransaksiRepository) TransaksiService {
	return &transaksiService{
		transaksiRepository: transaksiRep,
	}
}
