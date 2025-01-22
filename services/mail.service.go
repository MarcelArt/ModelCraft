package services

import (
	"github.com/MarcelArt/ModelCraft/config"
	"gopkg.in/gomail.v2"
)

type IMailService interface {
	SendMail(m Mailer) error
}

type MailService struct {
	dialer *gomail.Dialer
}

type Mailer struct {
	// From        string
	To          []string
	Subject     string
	Body        string
	Attachments []string
}

func NewMailService() *MailService {
	dialer := gomail.NewDialer(
		config.Env.SMTPHost,
		config.Env.SMTPPort,
		config.Env.SMTPEmail,
		config.Env.SMTPPassword,
	)

	return &MailService{
		dialer: dialer,
	}
}

func (s *MailService) SendMail(m Mailer) error {
	message := gomail.NewMessage()
	message.SetHeader("From", config.Env.SMTPName)
	message.SetHeader("To", m.To...)
	message.SetHeader("Subject", m.Subject)
	message.SetBody("text/html", m.Body)

	for _, attachment := range m.Attachments {
		message.Attach(attachment)
	}

	return s.dialer.DialAndSend(message)
}
