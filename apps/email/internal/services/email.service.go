package services

import (
	"context"
	"encoding/json"

	"github.com/go-gomail/gomail"
	"github.com/kitanoyoru/kita/apps/email/internal/config"
	"github.com/kitanoyoru/kita/apps/email/pkg/cache"
	pb "github.com/kitanoyoru/kita/apps/email/pkg/proto"
	"github.com/kitanoyoru/kita/apps/email/pkg/utils"
)

type CacheEmailLetter struct {
	Subject string `json:"subject"`
	Letter  string `json:"letter"`
}

func NewCacheEmailLetter(subject, letter string) *CacheEmailLetter {
	return &CacheEmailLetter{
		Subject: subject,
		Letter:  letter,
	}
}

type Email struct {
	creds *config.EmailConfig

	cache cache.Cache
}

func NewEmail() *Email {
	emailCreds, cacheCreds := config.Email(), config.Cache()

	cache := cache.NewRedis(cacheCreds.URL, cacheCreds.Password)

	return &Email{
		creds: emailCreds,
		cache: cache,
	}
}

func (e *Email) Init() error {
	htmlContent := utils.ReadFile(EmailLetterTemplatesPath + "confirmation-letter.html")
	cel := NewCacheEmailLetter(ConfirmationLetterSubject, htmlContent)

	data, err := json.Marshal(cel)
	if err != nil {
		return err
	}

	err = e.cache.Put(context.Background(), ConfirmationLetterKey, data)
	if err != nil {
		return err
	}

	return nil
}

func (e *Email) SendConfirmationMail(letterData *pb.SendOrderConfirmationRequest) error {
	data, err := e.cache.Get(context.Background(), ConfirmationLetterKey)
	if err != nil {
		return err
	}

	var cel CacheEmailLetter
	err = json.Unmarshal(data.([]byte), &cel)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()

	m.SetHeader("From", e.creds.SenderEmail)
	m.SetHeader("To", letterData.Email)
	m.SetHeader("Subject", cel.Subject)

	m.SetBody("text/html", cel.Letter)

	return e.sendLetter(m)
}

func (e *Email) sendLetter(letter *gomail.Message) error {
	d := gomail.NewDialer(e.creds.SMTPServer, e.creds.SMTPPort, e.creds.SenderEmail, e.creds.AppPassword)
	return d.DialAndSend(letter)
}
