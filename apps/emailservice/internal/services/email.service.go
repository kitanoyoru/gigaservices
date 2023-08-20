package services

import (
	"github.com/go-gomail/gomail"
	"github.com/kitanoyoru/kita/apps/emailservice/internal/config"
	pb "github.com/kitanoyoru/kita/apps/emailservice/pkg/proto"
	"github.com/kitanoyoru/kita/apps/emailservice/pkg/utils"
)

// REFACTOR: need to store subject and content of the letter in some hot memory in one structure

const (
	ConfirmationLetterSubject = "Kita cluster hosting confirmation"
)

type Email struct {
	creds *config.EmailConfig
}

func NewEmail() *Email {
	creds := config.Email()
	return &Email{
		creds,
	}
}

func (e *Email) SendConfirmationMail(letterData *pb.SendOrderConfirmationRequest) error {
	m := gomail.NewMessage()

	m.SetHeader("From", e.creds.SenderEmail)
	m.SetHeader("To", letterData.Email)
	m.SetHeader("Subject", ConfirmationLetterSubject)

	htmlContent := utils.ReadFile("./templates/letters/confirm.html")
	m.SetBody("text/html", htmlContent)

	return e.sendLetter(m)
}

func (e *Email) sendLetter(letter *gomail.Message) error {
	d := gomail.NewDialer(e.creds.SMTPServer, e.creds.SMTPPort, e.creds.SenderEmail, e.creds.AppPassword)
	return d.DialAndSend(letter)
}
