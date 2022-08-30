package sendgrid

import (
	"avatar/services/account_management/config"
	"avatar/services/account_management/domain/entity"
	"avatar/services/account_management/domain/mail"
	"bytes"
	"context"
	"embed"
	"fmt"
	"text/template"

	sendgrid "github.com/sendgrid/sendgrid-go"
	sendgrid_mail "github.com/sendgrid/sendgrid-go/helpers/mail"
)

//go:embed templates/*
var static embed.FS

type SendgridClient struct {
	client      *sendgrid.Client
	from        string
	frontendUri string
}

func NewMailClient(cfg *config.Environment) mail.MailClient {
	return &SendgridClient{
		client:      sendgrid.NewSendClient(cfg.SendgridApiKey),
		from:        cfg.MailFrom,
		frontendUri: cfg.FrontendUri,
	}
}

func (s *SendgridClient) CreateTemplateVariable(username string, token string) entity.ResetPassword {
	return entity.ResetPassword{
		Username:    username,
		ResetPwdURL: s.frontendUri,
		Token:       token,
	}
}

func (s *SendgridClient) GenerateTemplateByTemplateName(templateName string, bindVariable interface{}) (string, error) {
	templatePath := fmt.Sprintf("templates/%s", templateName)
	template, err := template.ParseFS(static, templatePath)
	if err != nil {
		return "", err
	}
	buf := &bytes.Buffer{}
	err = template.Execute(buf, bindVariable)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s\n", buf.String()), nil
}

func (s *SendgridClient) SendMail(
	ctx context.Context,
	subject string,
	toAddress string,
	fromName string,
	toName string,
	plaintextContent string,
	htmlContent string,
) error {
	from := sendgrid_mail.NewEmail(fromName, s.from)
	toEmail := sendgrid_mail.NewEmail(toName, toAddress)
	message := sendgrid_mail.NewSingleEmail(from, subject, toEmail, plaintextContent, htmlContent)
	_, err := s.client.Send(message)
	if err != nil {
		return err
	}
	return nil
}
