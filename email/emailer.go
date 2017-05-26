// Email contains a basic implementation of a SMTP email sender.  Usage :
//			em := email.NewEmailSender(email.EmailConfig{
//				ServerHostPort: "email-smtp.us-west-2.amazonaws.com:587",
//				Username: "AKIAI4CUKIYPFY6QB6IQ",
//				Password:"AiklA0mqVJnfJzctVHUImRduU5/8pqkUsEalj8Giz7Qe",
//			})
//
//			err := em.SendMail(&email.Email{
//				To:      "data@homes.co.nz",
//				From:    "greg.day@homes.co.nz",
//				Subject: "Test email",
//				Body:    "Test body",
//			})

package email

import (
	"bytes"
	"fmt"
	"net"
	"net/smtp"
	"time"

	"github.com/HomesNZ/go-common/env"
	"github.com/Sirupsen/logrus"
	"github.com/jordan-wright/email"
	"github.com/kr/pretty"
)

var (
	contextLogger = logrus.WithField("email", "sender")
	Emailer       Interface
)

type EmailConfig struct {
	Username       string
	Password       string
	ServerHostPort string
}

type EmailSender struct {
	Conf EmailConfig
	Send func(string, smtp.Auth, string, []string, []byte) error
}

func InitEmailer() {
	Emailer = NewEmailSender(EmailConfig{
		ServerHostPort: env.MustGetString("SMTP_HOST"),
		Username:       env.MustGetString("SMTP_USER"),
		Password:       env.MustGetString("SMTP_PASSWORD"),
	})
}

func NewEmailSender(conf EmailConfig) Interface {
	return &EmailSender{conf, smtp.SendMail}
}

// email.Interface is an interface to a struct that 'sends' emails. It allows the send function to be mocked out for testing.
type Interface interface {
	SendMail(email *Email) error
	SendMailWithAttachment(email *Email) error
}

func (e *EmailSender) createEmailBody(email *Email) (smtp.Auth, *bytes.Buffer) {
	host, _, _ := net.SplitHostPort(e.Conf.ServerHostPort)
	auth := smtp.PlainAuth("", e.Conf.Username, e.Conf.Password, host)
	pretty.Print("Hello")
	headers := make(map[string]string)
	headers["From"] = email.From
	headers["To"] = email.To
	headers["Subject"] = email.Subject
	headers["Date"] = time.Now().Format(time.UnixDate)

	message := bytes.NewBufferString("")
	for k, v := range headers {
		message.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}

	message.WriteString("\r\n")
	message.WriteString(email.Body)
	return auth, message
}

// Send sends a simple email via a smtp gateway using TLS
func (e *EmailSender) SendMail(email *Email) error {
	auth, message := e.createEmailBody(email)
	return e.Send(e.Conf.ServerHostPort, auth, email.From, []string{email.To}, message.Bytes())
}

func (sender *EmailSender) SendMailWithAttachment(emailContent *Email) error {
	auth, body := sender.createEmailBody(emailContent)
	e := &email.Email{
		To:      []string{emailContent.To},
		From:    emailContent.From,
		Subject: emailContent.Body,
		Text:    body.Bytes(),
	}
	e.AttachFile(emailContent.Attachment)
	return e.Send(sender.Conf.ServerHostPort, auth)
}
