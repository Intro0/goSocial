package mailer

import (
	"bytes"
	"fmt"
	"text/template"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridMailer struct {
	fromEmail string
	client    *sendgrid.Client
}

func NewSendGrid(apiKey, fromEmail string) *SendGridMailer {
	return &SendGridMailer{
		fromEmail: fromEmail,
		client:    sendgrid.NewSendClient(apiKey),
	}
}

func renderTemplate(templateFile string, data any) (subject string, body string, err error) {
	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return "", "", err
	}

	subjectBuffer := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(subjectBuffer, "subject", data); err != nil {
		return "", "", err
	}

	bodyBuffer := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(bodyBuffer, "body", data); err != nil {
		return "", "", err
	}

	return subjectBuffer.String(), bodyBuffer.String(), nil
}

func (m *SendGridMailer) buildMessage(templateFile, username, email string, data any, isSandbox bool) (*mail.SGMailV3, error) {
	subject, body, err := renderTemplate(templateFile, data)
	if err != nil {
		return nil, err
	}

	from := mail.NewEmail(FromName, m.fromEmail)
	to := mail.NewEmail(username, email)
	message := mail.NewSingleEmail(from, subject, to, "", body)
	message.SetMailSettings(&mail.MailSettings{
		SandboxMode: &mail.Setting{Enable: &isSandbox},
	})

	return message, nil
}

func (m *SendGridMailer) Send(templateFile, username, email string, data any, isSandbox bool) (int, error) {
	message, err := m.buildMessage(templateFile, username, email, data, isSandbox)
	if err != nil {
		return -1, err
	}

	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		response, err := m.client.Send(message)
		if err == nil {
			return response.StatusCode, nil
		}

		lastErr = err
		time.Sleep(time.Second * time.Duration(attempt+1))
	}

	return -1, fmt.Errorf("send email after %d attempts: %w", maxRetries, lastErr)
}
