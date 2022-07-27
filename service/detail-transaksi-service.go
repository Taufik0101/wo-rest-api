package service

import (
	"github.com/Taufik0101/wo-rest-api/dto"
	"github.com/Taufik0101/wo-rest-api/entity"
	"github.com/Taufik0101/wo-rest-api/repository"
	"github.com/mashingan/smapping"
)

type DetailTransaksiService interface {
	DetailTransaksiSeller(id_seller uint32) []entity.Detail_Transaksi
	UpdateDetailTransaksi(id_detail uint32, detail dto.UpdateDetailTransaksi) entity.Detail_Transaksi
	CreateDetailTransaksi(transaksi entity.Detail_Transaksi) entity.Detail_Transaksi
	DetailTransaksi(id_transaksi uint32) []entity.Detail_Transaksi
}

type detailTransaksiService struct {
	detailTransaksiRepository repository.DetailTransaksiRepository
}

func (d detailTransaksiService) DetailTransaksi(id_transaksi uint32) []entity.Detail_Transaksi {
	return d.detailTransaksiRepository.DetailTransaksi(id_transaksi)
}

func (d detailTransaksiService) CreateDetailTransaksi(transaksi entity.Detail_Transaksi) entity.Detail_Transaksi {
	newDetail := entity.Detail_Transaksi{}
	errSmap := smapping.FillStruct(&newDetail, smapping.MapFields(&transaksi))
	if errSmap != nil {
		panic("Gagal Parsing")
	}
	res := d.detailTransaksiRepository.CreateDetailTransaksi(newDetail)
	return res
}

func (d detailTransaksiService) DetailTransaksiSeller(id_seller uint32) []entity.Detail_Transaksi {
	return d.detailTransaksiRepository.DetailTransaksiSeller(id_seller)
}

func (d detailTransaksiService) UpdateDetailTransaksi(id_detail uint32, detail dto.UpdateDetailTransaksi) entity.Detail_Transaksi {
	upDetailTransaksi := entity.Detail_Transaksi{}
	errSmap := smapping.FillStruct(&upDetailTransaksi, smapping.MapFields(&detail))
	if errSmap != nil {
		panic("Gagal Parsing")
	}else {
		res := d.detailTransaksiRepository.UpdateDetailTransaksi(id_detail, upDetailTransaksi)
		return res
	}
}

func NewDetailTransaksiService(detailTransaksiRep repository.DetailTransaksiRepository) DetailTransaksiService {
	return &detailTransaksiService{
		detailTransaksiRepository: detailTransaksiRep,
	}
}
