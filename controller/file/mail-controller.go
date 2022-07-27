package file

import (
	"bytes"
	"fmt"
	"github.com/Taufik0101/wo-rest-api/dto"
	"gopkg.in/gomail.v2"
	"html/template"
	"os"
)

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		fmt.Println(err)
		return "", err
	}
	return buf.String(), nil
}

func ParseTemplateTolak(templateFileName string) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, nil); err != nil {
		fmt.Println(err)
		return "", err
	}
	return buf.String(), nil
}

func SendEmail(to string, subject string, data dto.Email, templateFile string) error {
	result, _ := ParseTemplate(templateFile, data)
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_USER"))
	m.SetHeader("To", to)
	// m.SetAddressHeader("Cc", "<RECIPIENT CC>", "<RECIPIENT CC NAME>")
	m.SetHeader("Subject", subject)
	//m.SetBody("text/html", "Email <b>"+to+"</b> and Password <i>"+data.Password+"</i>!")
	m.SetBody("text/html", result)
	//m.Attach(templateFile) // attach whatever you want
	senderPort := 587
	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), senderPort, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"))
	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}
	return err
}

func CreateEmailTolak(to string, subject string, templateFile string) error {
	result, _ := ParseTemplateTolak(templateFile)
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_USER"))
	m.SetHeader("To", to)
	// m.SetAddressHeader("Cc", "<RECIPIENT CC>", "<RECIPIENT CC NAME>")
	m.SetHeader("Subject", subject)
	//m.SetBody("text/html", "DITOLAK")
	m.SetBody("text/html", result)
	//m.Attach(templateFile) // attach whatever you want
	senderPort := 587
	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), senderPort, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"))
	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}
	return err
}

func SendEmailVerification(to string, data dto.Email) {
	var err error
	template := "./email/Account.html"
	subject := "noreply"
	err = SendEmail(to, subject, data, template)
	if err == nil {
		fmt.Println("send email '" + subject + "' success")
	} else {
		fmt.Println(err)
	}
}

func SendEmailTolak(to string)  {
	var err error
	template := "./email/tolak.html"
	subject := "noreply"
	err = CreateEmailTolak(to, subject, template)
	if err == nil {
		fmt.Println("send email '" + subject + "' success")
	} else {
		fmt.Println(err)
	}
}
