package account

import (
	"avatar/services/account_management/config"
	"avatar/services/account_management/domain/entity"
	"avatar/services/account_management/domain/repository"
	ctxUtil "avatar/services/account_management/pkg/ctx"
	stringUtil "avatar/services/account_management/pkg/string"
	pb "avatar/services/account_management/protos"
	"context"
	"errors"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func ResetPassword(
	ctx context.Context,
	db neo4j.Driver,
	accountRepository repository.AccountRepository,
	resetPasswordTokenRepository repository.ResetPasswordTokenRepository,
	input *pb.ResetPasswordRequest,
) (*pb.Empty, error) {
	passHashed, err := stringUtil.HashPass(input.NewPassword)
	if err != nil {
		return nil, errors.New("error hash pass")
	}
	input.NewPassword = passHashed
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		resetPasswordToken, err := ctxUtil.ExtractResetToken(ctx)
		if err != nil {
			return nil, err
		}
		resetPasswordTokenData, err := resetPasswordTokenRepository.GetByToken(ctx, resetPasswordToken)
		if err != nil {
			log.Error("error when write transaction, error: ", err)
			return nil, errors.New("invalid URL")
		}
		// check expire token
		if resetPasswordTokenData.ExpiresAt.Before(time.Now()) {
			err := resetPasswordTokenRepository.Delete(ctx, resetPasswordTokenData.Id)
			if err != nil {
				return nil, err
			}
			return nil, errors.New("reset password token is expired")
		}
		err = accountRepository.ChangePassword(ctx, &entity.ChangePasswordData{
			Id:          resetPasswordTokenData.AccountId,
			NewPassword: passHashed,
		})
		if err != nil {
			return nil, err
		}
		err = resetPasswordTokenRepository.Delete(ctx, resetPasswordTokenData.Id)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
