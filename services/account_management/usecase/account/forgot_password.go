package account

import (
	errUtil "avatar/pkg/err"
	"avatar/services/account_management/config"
	"avatar/services/account_management/domain/entity"
	"avatar/services/account_management/domain/mail"
	"avatar/services/account_management/domain/repository"
	stringUtil "avatar/services/account_management/pkg/string"
	pb "avatar/services/account_management/protos"
	"context"
	"errors"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

const (
	ResetPasswordTemplate = "reset_password.tmpl"
	ResetPasswordSubject  = "Reset password"
)

func ForgotPassword(
	ctx context.Context,
	cfg *config.Environment,
	db neo4j.Driver,
	accountRepository repository.AccountRepository,
	resetPasswordTokenRepository repository.ResetPasswordTokenRepository,
	mailClient mail.MailClient,
	input *pb.ForgotPasswordRequest,
) (*pb.Empty, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	resetPasswordTokenData, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		account, err := accountRepository.GetByEmail(ctx, input.Email)
		if err != nil {
			return nil, err
		}
		tokenData := &entity.ResetPasswordToken{
			AccountId: account.Id,
			Email:     account.Email,
			Username:  account.Username,
			Token:     stringUtil.RandString(10),
			ExpiresAt: time.Now().UTC().Add(time.Minute * time.Duration(cfg.ResetTokenExpirationMinute)),
		}
		token, err := resetPasswordTokenRepository.Create(ctx, tokenData)
		if err != nil {
			return nil, err
		}
		return token, nil
	})
	if err != nil {
		if err.Error() == errUtil.ERR_NO_RECORD {
			return nil, errors.New("email does not exist")
		}
		return nil, err
	}
	resetToken := resetPasswordTokenData.(*entity.ResetPasswordToken)
	resetPasswordData := mailClient.CreateTemplateVariable(resetToken.Username, resetToken.Token)
	template, err := mailClient.GenerateTemplateByTemplateName(ResetPasswordTemplate, resetPasswordData)
	if err != nil {
		log.Error("Error generate template sendgrid! ", err)
		return nil, err
	}
	err = mailClient.SendMail(ctx, ResetPasswordSubject, input.Email, "", "", template, template)
	if err != nil {
		log.Error("Error sendgrid send mail! ", err)
		return nil, err
	}
	return &pb.Empty{}, nil
}
