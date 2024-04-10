package mail

import (
	"sync"

	"github.com/spf13/viper"
)

type From struct{
	Address string
	Name string
}

type Email struct{
	From From
	To []string
	Bcc []string
	Cc []string
	Subject string
	Text []byte
	HTML []byte
}

type Mailer struct{
	Driver Driver
}

var once sync.Once
var internalMailer *Mailer

func NewMailer()*Mailer{
	once.Do(func() {
		internalMailer=&Mailer{
			Driver: &SMTP{},
		}
	})

	return internalMailer
}

func(mail *Mailer)Send(email Email)bool{
	return mail.Driver.Send(email,viper.GetStringMapString("mail.smtp"))
}