package service

import (
	"github.com/Taufik0101/wo-rest-api/dto"
	"github.com/Taufik0101/wo-rest-api/entity"
	"github.com/Taufik0101/wo-rest-api/repository"
	"github.com/mashingan/smapping"
)

type DetailAlbumService interface {
	AllDetailAlbum() []entity.Detail_Album
	DetailAlbumByAlbum(id_album uint32) []entity.Detail_Album
	SimpanDetailAlbum(detailAl dto.CreateDetailAlbum) entity.Detail_Album
	HapusDetail(id_detailAl uint32)
	DetailById(id_detail uint32) entity.Detail_Album
}

type detailAlbumService struct {
	detailAlbumRepository repository.DetailAlbumRepository
}

func (d detailAlbumService) DetailById(id_detail uint32) entity.Detail_Album {
	return d.detailAlbumRepository.AlbumById(id_detail)
}

func (d detailAlbumService) AllDetailAlbum() []entity.Detail_Album {
	return d.detailAlbumRepository.AllDetailAlbum()
}

func (d detailAlbumService) DetailAlbumByAlbum(id_album uint32) []entity.Detail_Album {
	return d.detailAlbumRepository.DetailAlbumByAlbum(id_album)
}

func (d detailAlbumService) SimpanDetailAlbum(detailAl dto.CreateDetailAlbum) entity.Detail_Album {
	newDetail := entity.Detail_Album{}
	errSmap := smapping.FillStruct(&newDetail, smapping.MapFields(&detailAl))
	if errSmap != nil {
		panic("Gagal Parsing")
	}
	res := d.detailAlbumRepository.SimpanDetailAlbum(newDetail)
	return res
}

func (d detailAlbumService) HapusDetail(id_detailAl uint32) {
	d.detailAlbumRepository.HapusDetailAl(id_detailAl)
}

func NewDetailAlbumService(detailAlbumRep repository.DetailAlbumRepository) DetailAlbumService {
	return &detailAlbumService{
		detailAlbumRepository: detailAlbumRep,
	}
}