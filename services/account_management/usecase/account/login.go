package account

import (
	errUtil "avatar/pkg/err"
	"avatar/pkg/jwt"
	"avatar/services/account_management/config"
	"avatar/services/account_management/domain/entity"
	"avatar/services/account_management/domain/repository"
	stringUtil "avatar/services/account_management/pkg/string"
	pb "avatar/services/account_management/protos"
	"context"
	"errors"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func Login(
	ctx context.Context,
	cfg *config.Environment,
	db neo4j.Driver,
	accountRepository repository.AccountRepository,
	input *pb.LoginRequest,
) (*pb.LoginResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	accountRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		acc, err := accountRepository.GetByEmail(ctx, input.Email)
		if err != nil {
			return nil, err
		}
		return acc, nil
	})
	if err != nil {
		if err.Error() == errUtil.ERR_NO_RECORD {
			return nil, errors.New("wrong email or password")
		}
		return nil, err
	}
	account := accountRaw.(*entity.Account)
	hashedPassword := account.Password
	if err := stringUtil.CheckPasswordHash(hashedPassword, input.Password); err != nil {
		return nil, errors.New("wrong email or password")
	}
	token, err := jwt.CreateJWT(account, cfg.JwtSecretKey, cfg.JwtExpirationHour)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		Token:       token,
		UserId:      account.Id,
		DisplayName: account.Username,
		Avatar:      account.Avatar,
		Gender:      account.Gender,
	}, nil
}
