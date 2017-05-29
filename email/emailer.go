package email

import (
	"net"
	"net/smtp"
	"sync"

	"github.com/HomesNZ/go-common/env"
	"github.com/Sirupsen/logrus"
	"github.com/jordan-wright/email"
)

var (
	contextLogger = logrus.WithField("email", "sender")
	once          sync.Once
	Emailer       EmailSender
)

type EmailSender struct {
	Username       string
	Password       string
	ServerHostPort string
}

func InitEmailer() {
	once.Do(func() {
		Emailer = EmailSender{
			ServerHostPort: env.MustGetString("SMTP_HOST"),
			Username:       env.MustGetString("SMTP_USER"),
			Password:       env.MustGetString("SMTP_PASSWORD"),
		}
	})
}

// Send sends a simple email via a smtp gateway using TLS
func Send(content *Email) error {
	InitEmailer()
	host, _, _ := net.SplitHostPort(Emailer.ServerHostPort)

	auth := smtp.PlainAuth("", Emailer.Username, Emailer.Password, host)

	eml := &email.Email{
		To:      []string{content.To},
		From:    content.From,
		Subject: content.Subject,
		Text:    []byte(content.Body),
	}

	if content.Attachment != "" {
		_, err := eml.AttachFile(content.Attachment)
		if err != nil {
			return err
		}
	}
	return eml.Send(Emailer.ServerHostPort, auth)
}
