package email

import (
	"chatgpt-web-new-go/common/config"
	"gopkg.in/gomail.v2"
)

func InitEmailDialer() {
	server := config.Config.EmailServer

	config.EmailDialer = gomail.NewDialer(server.Host, server.Port, server.User, server.Password)
}

func SendMail(subject, toAddress, content string) error {
	server := config.Config.EmailServer

	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(server.User, server.SenderName))
	m.SetHeader("To", toAddress)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)
	err := config.EmailDialer.DialAndSend(m)
	return err
}
