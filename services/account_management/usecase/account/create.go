package account

import (
	errUtil "avatar/pkg/err"
	"avatar/services/account_management/config"
	"avatar/services/account_management/domain/entity"
	"avatar/services/account_management/domain/mail"
	"avatar/services/account_management/domain/repository"
	ctxUtil "avatar/services/account_management/pkg/ctx"
	stringUtil "avatar/services/account_management/pkg/string"
	pb "avatar/services/account_management/protos"
	"context"
	"errors"
	"time"

	"github.com/jinzhu/copier"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func Create(
	ctx context.Context,
	cfg *config.Environment,
	db neo4j.Driver,
	accountRepository repository.AccountRepository,
	resetPasswordTokenRepository repository.ResetPasswordTokenRepository,
	mailClient mail.MailClient,
	input *pb.CreateAccountRequest,
) (*pb.Account, error) {
	passHashed, err := stringUtil.HashPass(stringUtil.RandString(10))
	if err != nil {
		return nil, errors.New("error hash pass")
	}
	data := entity.Account{Password: passHashed}
	err = copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	// get uid admin for detect createdBy
	adminId, err := ctxUtil.ExtractUid(ctx)
	if err != nil {
		return nil, err
	}
	data.CreatedBy = adminId
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	newAcc, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		_, err := accountRepository.GetByEmail(ctx, input.Email)
		if err == nil {
			return nil, errors.New("email is taken")
		}
		if err.Error() != errUtil.ERR_NO_RECORD {
			return nil, err
		}
		acc, err := accountRepository.Create(ctx, &data)
		if err != nil {
			return nil, err
		}
		tokenData := &entity.ResetPasswordToken{
			AccountId: acc.Id,
			Email:     acc.Email,
			Username:  acc.Username,
			Token:     stringUtil.RandString(10),
			ExpiresAt: time.Now().UTC().Add(time.Minute * time.Duration(cfg.ResetTokenExpirationMinute)),
		}
		token, err := resetPasswordTokenRepository.Create(ctx, tokenData)
		if err != nil {
			return nil, err
		}
		// sendgrid send mail
		resetPasswordData := mailClient.CreateTemplateVariable(tokenData.Username, token.Token)
		template, err := mailClient.GenerateTemplateByTemplateName(ResetPasswordTemplate, resetPasswordData)
		if err != nil {
			return nil, err
		}
		err = mailClient.SendMail(ctx, ResetPasswordSubject, input.Email, "", "", template, template)
		if err != nil {
			return nil, err
		}
		return acc, nil
	})
	if err != nil {
		return nil, err
	}
	account := newAcc.(*entity.Account)
	var result pb.Account
	err = copier.Copy(&result, account)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
