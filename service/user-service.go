package service

import (
	"github.com/Taufik0101/wo-rest-api/dto"
	"github.com/Taufik0101/wo-rest-api/entity"
	"github.com/Taufik0101/wo-rest-api/repository"
	"github.com/mashingan/smapping"
)

type UserService interface {
	ListAll() []entity.User
	FindByID(id_user uint32) entity.User
	FindByEmail(email string) interface{}
	VendorByCategory(caetgory uint32) []entity.User
	VendorByCity(kota uint32) []entity.User
	VendorByKategoryAndCity(category uint32, city uint32) []entity.User
	AllVendor() []entity.User
	UpdateUser(id_user uint32, user dto.Update) entity.User
	CekPassword(user_id uint32, password string) bool
}

type userService struct {
	userRepository repository.UserRepository
}

func (u userService) CekPassword(user_id uint32, password string) bool {
	user := u.userRepository.FindByID(user_id)
	cekPas := repository.ComparePassword(user.Password, password)
	if cekPas {
		return true
	}
	return false
}

func (u userService) ListAll() []entity.User {
	return u.userRepository.ListAll()
}

func (u userService) FindByID(id_user uint32) entity.User {
	return u.userRepository.FindByID(id_user)
}

func (u userService) FindByEmail(email string) interface{} {
	return u.userRepository.FindByEmail(email)
}

func (u userService) VendorByCategory(caetgory uint32) []entity.User {
	return u.userRepository.VendorByCategory(caetgory)
}

func (u userService) VendorByCity(kota uint32) []entity.User {
	return u.userRepository.VendorByCity(kota)
}

func (u userService) VendorByKategoryAndCity(category uint32, city uint32) []entity.User {
	return u.userRepository.VendorByKategoryAndCity(category, city)
}

func (u userService) AllVendor() []entity.User {
	return u.userRepository.AllVendor()
}

func (u userService) UpdateUser(id_user uint32, user dto.Update) entity.User {
	upUser := entity.User{}
	errSmap := smapping.FillStruct(&upUser, smapping.MapFields(&user))
	if errSmap != nil {
		panic("Gagal Parsing")
	}else {
		res := u.userRepository.UpdateUser(id_user, upUser)
		return res
	}
}

func NewUserService(userRep repository.UserRepository) UserService {
	return &userService{
		userRepository: userRep,
	}
}
