package email_test

import (
	"github.com/HomesNZ/data-import/email"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/smtp"
)

func NewMockEmailSender(conf email.EmailConfig) (email.Interface, *emailRecorder) {
	r := emailRecorder{}
	f := func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		r = emailRecorder{addr: addr, auth: a, from: from, to: to, msg: msg}
		return nil
	}

	return &email.EmailSender{conf, f}, &r
}

type emailRecorder struct {
	addr string
	auth smtp.Auth
	from string
	to   []string
	msg  []byte
}

var _ = Describe("Email Test", func() {
	Context("Valid email", func() {
		It("Valid login", func() {
			sender, recorder := NewMockEmailSender(email.EmailConfig{})
			sender.SendMail(&email.Email{
				To:      "data@homes.co.nz",
				From:    "greg.day@homes.co.nz",
				Subject: "Test email",
				Body:    "Test body",
			})

			Expect(recorder.from).To(Equal("greg.day@homes.co.nz"))
			Expect(recorder.to).To(Equal([]string{"data@homes.co.nz"}))
			Expect(string(recorder.msg)).To(MatchRegexp(`.*Test body.*`))
		})
	})
})
