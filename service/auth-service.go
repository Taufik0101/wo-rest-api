package service

import (
	"bytes"
	"fmt"
	"github.com/Taufik0101/wo-rest-api/dto"
	"github.com/Taufik0101/wo-rest-api/entity"
	"github.com/Taufik0101/wo-rest-api/repository"
	"github.com/mashingan/smapping"
	"html/template"
	"net/smtp"
)

type AuthService interface {
	VerifyPassword(email string, password string) interface{}
	Register(regis dto.Register) entity.User
	IsDuplicateEmail(email string) bool
}

type authservice struct {
	authRepository repository.AuthRepository
}

func (a authservice) VerifyPassword(email string, password string) interface{} {
	return a.authRepository.VerifyPassword(email, password)
}

func (a authservice) Register(regis dto.Register) entity.User {
	newUser := entity.User{}
	errSmap := smapping.FillStruct(&newUser, smapping.MapFields(&regis))
	if errSmap != nil {
		panic("Gagal Parsing")
	}
	result := a.authRepository.Register(newUser)
	return result
}

func (a authservice) IsDuplicateEmail(email string) bool {
	res := a.authRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func NewAuthService(authRep repository.AuthRepository) AuthService {
	return &authservice{
		authRepository: authRep,
	}
}

func SendMail(email string, password string) bool {
	from := "dreamswedding06@gmail.com"
	pass := "dreamsbangetbuatlulus64:_x"

	// Receiver email address.
	to := []string{
		"hid09h@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, pass, smtpHost)

	t, _ := template.ParseFiles("Account.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Email    string
		Password string
	}{
		Email:    email,
		Password: password,
	})

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("Email Sent!")
	return true
}
