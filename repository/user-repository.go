package repository

import (
	"github.com/Taufik0101/wo-rest-api/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	ListAll() []entity.User
	FindByID(id_user uint32) entity.User
	FindByEmail(email string) interface{}
	VendorByCategory(caetgory uint32) []entity.User
	VendorByCity(kota uint32) []entity.User
	VendorByKategoryAndCity(category uint32, city uint32) []entity.User
	AllVendor() []entity.User
	UpdateUser(id_user uint32, user entity.User) entity.User
}

type userRepository struct {
	connection *gorm.DB
}

func (u userRepository) UpdateUser(id_user uint32, user entity.User) entity.User {
	u.connection.Where("id = ?", id_user).Updates(&user)
	u.connection.Where("id = ?", id_user).Preload("Cat").Preload("Kot").Find(&user)
	return user
}

func (u userRepository) AllVendor() []entity.User {
	var vendors []entity.User
	u.connection.Where("role = 'Vendor'").Preload("Cat").Preload("Kot").Find(&vendors)
	return vendors
}

func (u userRepository) VendorByCity(kota uint32) []entity.User {
	var vendors []entity.User
	u.connection.Where("kota = ?", kota).Preload("Cat").Preload("Kot").Find(&vendors)
	return vendors
}

func (u userRepository) VendorByKategoryAndCity(category uint32, city uint32) []entity.User {
	var vendors []entity.User
	u.connection.Where("category = ? AND kota = ?", category, city).Preload("Cat").Preload("Kot").Find(&vendors)
	return vendors
}

func (u userRepository) VendorByCategory(category uint32) []entity.User {
	var vendors []entity.User
	u.connection.Where("category = ?", category).Preload("Cat").Preload("Kot").Find(&vendors)
	return vendors
}

func (u userRepository) ListAll() []entity.User {
	var users []entity.User
	u.connection.Order("id asc").Preload("Cat").Preload("Kot").Find(&users)
	return users
}

func (u userRepository) FindByID(id_user uint32) entity.User {
	var user entity.User
	u.connection.Where("id = ?", id_user).Preload("Cat").Preload("Kot").Take(&user)
	return user
}

func (u userRepository) FindByEmail(email string) interface{} {
	var user entity.User
	res := u.connection.Where("email = ?", email).Preload("Cat").Preload("Kot").Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func NewUserRepository(conn *gorm.DB) UserRepository {
	return &userRepository{
		connection: conn,
	}
}
