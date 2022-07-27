package service

import (
	"github.com/Taufik0101/wo-rest-api/dto"
	"github.com/Taufik0101/wo-rest-api/entity"
	"github.com/Taufik0101/wo-rest-api/repository"
	"github.com/golang-module/carbon"
)

type AlbumService interface {
	AllAlbum() []entity.Album
	SimpanAlbum(album dto.CreateAlbum) entity.Album
	AlbumByVendor(id_vendor uint32) []entity.Album
	UpdateAlbum(id_album uint32, album dto.UpdateAlbum) entity.Album
}

type albumService struct {
	albumRepository repository.AlbumRepository
}

func (a albumService) UpdateAlbum(id_album uint32, album dto.UpdateAlbum) entity.Album {
	upAlbum := entity.Album{
		Nama:                album.Nama,
		Tanggal_Pelaksanaan: carbon.Parse(album.Tanggal_Pelaksanaan, "Asia/Jakarta").Time,
	}
	res := a.albumRepository.UpdateAlbum(id_album, upAlbum)
	return res
}

func (a albumService) AllAlbum() []entity.Album {
	return a.albumRepository.AllAlbum()
}

func (a albumService) SimpanAlbum(album dto.CreateAlbum) entity.Album {
	newAlbum := entity.Album{
		IdVendor:            album.IdVendor,
		Nama:                album.Nama,
		Tanggal_Pelaksanaan: carbon.Parse(album.Tanggal_Pelaksanaan, "Asia/Jakarta").Time,
	}
	res := a.albumRepository.SimpanAlbum(newAlbum)
	return res
}

func (a albumService) AlbumByVendor(id_vendor uint32) []entity.Album {
	return a.albumRepository.AlbumByVendor(id_vendor)
}

func NewAlbumService(albumRep repository.AlbumRepository) AlbumService {
	return &albumService{
		albumRepository: albumRep,
	}
}
