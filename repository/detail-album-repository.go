package repository

import (
	"github.com/Taufik0101/wo-rest-api/entity"
	"gorm.io/gorm"
)

type DetailAlbumRepository interface {
	AllDetailAlbum() []entity.Detail_Album
	DetailAlbumByAlbum(id_album uint32) []entity.Detail_Album
	SimpanDetailAlbum(detailAl entity.Detail_Album) entity.Detail_Album
	HapusDetailAl(id_detailAl uint32)
	AlbumById(id_detail uint32) entity.Detail_Album
}

type detailAlbumRepository struct {
	connection *gorm.DB
}

func (d detailAlbumRepository) AlbumById(id_detail uint32) entity.Detail_Album {
	var detailalbums entity.Detail_Album
	d.connection.Where("id_detail = ?", id_detail).Preload("Albums").Find(&detailalbums)
	return detailalbums
}

func (d detailAlbumRepository) AllDetailAlbum() []entity.Detail_Album {
	var detailalbums []entity.Detail_Album
	d.connection.Order("id_detail asc").Preload("Albums").Find(&detailalbums)
	return detailalbums
}

func (d detailAlbumRepository) DetailAlbumByAlbum(id_album uint32) []entity.Detail_Album {
	var detailalbums []entity.Detail_Album
	d.connection.Where("id_album = ?", id_album).Preload("Albums").Find(&detailalbums)
	return detailalbums
}

func (d detailAlbumRepository) SimpanDetailAlbum(detailAl entity.Detail_Album) entity.Detail_Album {
	d.connection.Save(&detailAl)
	d.connection.Preload("Albums").Find(&detailAl)
	return detailAl
}

func (d detailAlbumRepository) HapusDetailAl(id_detailAl uint32) {
	d.connection.Where("id_detail = ?", id_detailAl).Delete(&entity.Detail_Album{})
}

func NewDetailAlbumRepository(conn *gorm.DB) DetailAlbumRepository {
	return &detailAlbumRepository{
		connection: conn,
	}
}
