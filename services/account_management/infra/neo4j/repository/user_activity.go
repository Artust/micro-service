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

type userActivityRepositoryImpl struct {
	baseRepository base_repository.BaseRepository
}

func NewUserActivityRepository(
	baseRepository base_repository.BaseRepository,
) repository.UserActivityRepository {
	return &userActivityRepositoryImpl{
		baseRepository: baseRepository,
	}
}

func (r *userActivityRepositoryImpl) Get(ctx context.Context, accountId int64) ([]*entity.UserActivity, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	result := []*entity.UserActivity{}
	err := r.baseRepository.Model(&entity.UserActivity{}).Where("accountId = ?", accountId).Find(transaction, &result)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return result, nil
}

func (r *userActivityRepositoryImpl) Save(ctx context.Context, input *entity.UserActivity) (*entity.UserActivity, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Create(transaction, input)
	if err != nil {
		log.Error("error when write transaction save activity user, error: ", err)
		return nil, err
	}
	return input, nil
}
