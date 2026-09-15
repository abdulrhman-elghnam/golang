package email

import (
	"os"

	"github.com/resend/resend-go/v3"
)

func SendEmail(
	from string,
	to string,
	subject string,
	html string,
) error {
	client := resend.NewClient(os.Getenv("RESEND_KEY"))

	params := &resend.SendEmailRequest{
		From:    from,
		To:      []string{to},
		Subject: subject,
		Html:    html,
	}

	_, err := client.Emails.Send(params)

	return err
}