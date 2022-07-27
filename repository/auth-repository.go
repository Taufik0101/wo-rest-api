package repository

import (
	"github.com/Taufik0101/wo-rest-api/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user entity.User) entity.User
	VerifyPassword(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
}

type authRepository struct {
	connection *gorm.DB
	userRepo UserRepository
}

func (a authRepository) Register(user entity.User) entity.User {
	hashPW, _ := Hash(user.Password)
	user.Password = string(hashPW)
	a.connection.Save(&user)
	a.connection.Preload("Cat").Preload("Kot").Find(&user)
	return user
}

func (a authRepository) VerifyPassword(email string, password string) interface{} {
	findEmail := a.userRepo.FindByEmail(email)
	if value, ok := findEmail.(entity.User); ok {
		checkPass := ComparePassword(value.Password, password)
		if checkPass {
			return findEmail
		}
		return false
	}
	return false
}

func (a authRepository) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return a.connection.Where("email = ?", email).Take(&user)
}

func NewAuthRepository(conn *gorm.DB, userRep UserRepository) AuthRepository {
	return &authRepository{
		connection: conn,
		userRepo: userRep,
	}
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}