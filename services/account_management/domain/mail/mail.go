package mail

import (
	"avatar/services/account_management/domain/entity"
	"context"
)

type MailClient interface {
	CreateTemplateVariable(username string, token string) entity.ResetPassword
	GenerateTemplateByTemplateName(templateName string, bindVariable interface{}) (string, error)
	SendMail(ctx context.Context, subject, toAddress, fromName string, toName string, plaintextContent, htmlContent string) error
}
