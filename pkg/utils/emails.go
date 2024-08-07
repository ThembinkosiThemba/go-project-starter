package utils

import (
	"fmt"
	"os"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailService struct {
	SenderName  string
	SenderEmail string
	APIKey      string
}

func NewEmailService() *EmailService {
	return &EmailService{
		SenderName:  os.Getenv("ADMIN_NAME"),
		SenderEmail: os.Getenv("ADMIN_EMAIL"),
		APIKey:      os.Getenv("SENDGRID_API_KEY"),
	}
}

func (es *EmailService) SendEmail(toUser, toEmail string) error {
	from := mail.NewEmail(es.SenderName, es.SenderEmail)
	to := mail.NewEmail(toUser, toEmail)

	subject := ""
	plaintTextContent := ""
	htmlContent := ""

	message := mail.NewSingleEmail(from, subject, to, plaintTextContent, htmlContent)
	client := sendgrid.NewSendClient(es.APIKey)

	response, err := client.Send(message)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	if response.StatusCode >= 300 {
		return fmt.Errorf("unexpected status code: %d, body: %s", response.StatusCode, response.Body)
	}
	return nil

}
