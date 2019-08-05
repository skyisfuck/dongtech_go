package handler

import (
	"dongtech_go/config"
	"github.com/go-gomail/gomail"
	"github.com/sirupsen/logrus"
)

func SendEmail(from string, to string, cc string, subject string, body string, attachment string) error {
	m := gomail.NewMessage()
	var err error
	if from != "" {
		m.SetHeader("From", from)
	}
	if to != "" {
		m.SetHeader("To", to)
	}
	if cc != "" {
		m.SetAddressHeader("Cc", cc, "")
	}
	if subject != "" {
		m.SetHeader("Subject", subject)
	}
	if body != "" {
		m.SetBody("text/html", body)
	}
	if from != "" {
		m.Attach(attachment)
	}

	config, err := config.GetConfig()
	if err != nil {
		return err
	}

	if config.Email != nil {
		d := gomail.NewDialer(config.Email.Host, config.Email.Port, config.Email.Username, config.Email.Password)
		// Send the email to Bob, Cora and Dan.
		if err = d.DialAndSend(m); err != nil {
			logrus.WithError(err).WithField("email", config.Email).Println("send email err")
			panic(err)
		}
	} else {
		logrus.Println("get email config err")
	}
	return nil
}
