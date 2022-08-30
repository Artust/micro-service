package repository

import (
	"avatar/pkg/base_repository"
	"avatar/services/account_management/config"
	"avatar/services/account_management/domain/entity"
	"avatar/services/account_management/domain/repository"
	"context"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

type resetPasswordTokenRepositoryImpl struct {
	baseRepository base_repository.BaseRepository
}

func NewResetPasswordTokenRepository(
	baseRepository base_repository.BaseRepository,
) repository.ResetPasswordTokenRepository {
	return &resetPasswordTokenRepositoryImpl{
		baseRepository: baseRepository,
	}
}

func (r *resetPasswordTokenRepositoryImpl) Create(ctx context.Context, input *entity.ResetPasswordToken) (*entity.ResetPasswordToken, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Create(transaction, input)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *resetPasswordTokenRepositoryImpl) GetByToken(ctx context.Context, token string) (*entity.ResetPasswordToken, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	var result entity.ResetPasswordToken
	err := r.baseRepository.Model(&entity.ResetPasswordToken{}).Where("token = ?", token).FindOne(transaction, &result)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return &result, nil
}

func (r *resetPasswordTokenRepositoryImpl) Delete(ctx context.Context, id int64) (err error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	records, err := transaction.Run(`MATCH (r:ResetPasswordToken) WHERE id(r) = $id DELETE r`, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		log.Error("error when delete routine center, error: ", err)
		return
	}
	_, err = records.Consume()
	if err != nil {
		log.Error("error when record delete routine center, error: ", err)
		return
	}
	return
}
