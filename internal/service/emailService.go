package service

import (
	"log"
	"net/http"
	"net/smtp"
)

type EmailService struct {
	client           *http.Client
	emailApiUsername string
	emailApiPassword string
}

func NewEmailService(emailApiUsername string, emailApiPassword string) *EmailService {
	return &EmailService{
		client:           &http.Client{},
		emailApiUsername: emailApiUsername,
		emailApiPassword: emailApiPassword,
	}
}

func (es *EmailService) sendMesage(message string, email string) error {
	from := "dlarin-auth-service@itis.ru"
	to := email
	subject := "Auth email code"
	body := message

	smtpHost := "sandbox.smtp.mailtrap.io"
	smtpPort := "587"

	auth := smtp.PlainAuth("", es.emailApiUsername, es.emailApiPassword, smtpHost)

	request := []byte("Subject: " + subject + "\r\n" +
		"From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, request)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Email sent successfully!")
	return nil
}
