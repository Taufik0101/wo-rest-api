package repository

import (
	"github.com/Taufik0101/wo-rest-api/entity"
	"gorm.io/gorm"
)

type KotaRepository interface {
	AllKota() []entity.Kota
	SimpanKota(kota entity.Kota) entity.Kota
	UpdateKota(id_kota uint32, kota entity.Kota) entity.Kota
	DeleteKota(id_kota uint32)
	FindByName(nama string) (tx *gorm.DB)
}

type kotaRepository struct {
	connection *gorm.DB
}

func (k kotaRepository) FindByName(nama string) (tx *gorm.DB) {
	var kota entity.Kota
	return k.connection.Where("nama_kota = ?", nama).Take(&kota)
}

func (k kotaRepository) AllKota() []entity.Kota {
	var kotas []entity.Kota
	k.connection.Order("id_kota asc").Find(&kotas)
	return kotas
}

func (k kotaRepository) SimpanKota(kota entity.Kota) entity.Kota {
	k.connection.Save(&kota)
	return kota
}

func (k kotaRepository) UpdateKota(id_kota uint32, kota entity.Kota) entity.Kota {
	k.connection.Where("id_kota = ?", id_kota).Updates(&kota)
	return kota
}

func (k kotaRepository) DeleteKota(id_kota uint32) {
	k.connection.Where("id_kota = ?", id_kota).Delete(&entity.Kota{})
}

func NewKotaRepository(conn *gorm.DB) KotaRepository {
	return &kotaRepository{
		connection: conn,
	}
}