package notify

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

type Service interface {
	VerifyEmail(emailId string) error
}

type service struct{}

func NewNotifyService() Service {
	return &service{}
}

func (s *service) VerifyEmail(emailId string) error {
	mail := &email.Email{
		To:      []string{"test@example.com"},
		From:    "Rijin Elayambari <test@gmail.com>",
		Subject: "Verify Email",
		Text:    []byte("Click confirmation url to verify the email address"),
		HTML: []byte("<h1>Verify</h1>" +
			"<span>Click confirmation url to verify the email address</span>"),
	}

	err := mail.Send("smtp.gmail.com:587", smtp.PlainAuth("", "test@gmail.com", "password123", "smtp.gmail.com"))
	return err
}
