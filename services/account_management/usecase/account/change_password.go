package account

import (
	"avatar/services/account_management/config"
	"avatar/services/account_management/domain/entity"
	"avatar/services/account_management/domain/repository"
	stringUtil "avatar/services/account_management/pkg/string"
	pb "avatar/services/account_management/protos"
	"context"
	"errors"

	"github.com/jinzhu/copier"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func ChangePassword(
	ctx context.Context,
	db neo4j.Driver,
	accountRepository repository.AccountRepository,
	input *pb.ChangePasswordRequest,
) (*pb.Empty, error) {
	data := entity.ChangePasswordData{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	passHashed, err := stringUtil.HashPass(input.NewPassword)
	if err != nil {
		return nil, errors.New("error hash pass")
	}
	data.NewPassword = passHashed
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		acc, err := accountRepository.GetById(ctx, input.Id)
		if err != nil {
			return nil, err
		}
		hashedPassword := acc.Password
		if err := stringUtil.CheckPasswordHash(hashedPassword, input.Password); err != nil {
			return nil, errors.New("wrong email or password")
		}
		err = accountRepository.ChangePassword(ctx, &data)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, err
}
