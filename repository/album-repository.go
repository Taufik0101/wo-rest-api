package repository

import (
	"github.com/Taufik0101/wo-rest-api/entity"
	"gorm.io/gorm"
)

type AlbumRepository interface {
	AllAlbum() []entity.Album
	SimpanAlbum(album entity.Album) entity.Album
	AlbumByVendor(id_vendor uint32) []entity.Album
	UpdateAlbum(id_album uint32, album entity.Album) entity.Album
}

type albumRepository struct {
	connection *gorm.DB
}

func (a albumRepository) UpdateAlbum(id_album uint32, album entity.Album) entity.Album {
	a.connection.Where("id_album = ?", id_album).Updates(&album)
	a.connection.Where("id_album = ?", id_album).Preload("Vendor").Preload("Detail").Find(&album)
	return album
}

func (a albumRepository) AllAlbum() []entity.Album {
	var albums []entity.Album
	a.connection.Order("id_album asc").Preload("Vendor").Preload("Detail").Find(&albums)
	return albums
}

func (a albumRepository) SimpanAlbum(album entity.Album) entity.Album {
	a.connection.Save(&album)
	a.connection.Preload("Vendor").Preload("Detail").Find(&album)
	return album
}

func (a albumRepository) AlbumByVendor(id_vendor uint32) []entity.Album {
	var albums []entity.Album
	a.connection.Where("id_vendor = ?", id_vendor).Preload("Vendor").Preload("Detail").Find(&albums)
	return albums
}

func NewAlbumRepository(conn *gorm.DB) AlbumRepository {
	return &albumRepository{
		connection: conn,
	}
}
