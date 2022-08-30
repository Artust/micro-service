package repository

import (
	"avatar/pkg/base_repository"
	errUtil "avatar/pkg/err"
	"avatar/services/account_management/config"
	"avatar/services/account_management/domain/entity"
	"avatar/services/account_management/domain/repository"
	"context"
	"errors"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

type accountRoleRepositoryImpl struct {
	baseRepository base_repository.BaseRepository
}

func NewAccountRoleRepository(
	baseRepository base_repository.BaseRepository,
) repository.AccountRoleRepository {
	return &accountRoleRepositoryImpl{
		baseRepository: baseRepository,
	}
}

func (r *accountRoleRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.AccountRole, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	var result entity.AccountRole
	err := r.baseRepository.Model(&entity.AccountRole{}).Where("id = ?", id).FindOne(transaction, &result)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return &result, nil
}

func (r *accountRoleRepositoryImpl) GetList(ctx context.Context) ([]*entity.AccountRole, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	result := []*entity.AccountRole{}
	err := r.baseRepository.Model(&entity.AccountRole{}).Find(transaction, &result)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return result, nil
}

func (r *accountRoleRepositoryImpl) Create(ctx context.Context, role *entity.AccountRole) (*entity.AccountRole, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Create(transaction, role)
	if err != nil {
		log.Error("error when write transaction create role, error: ", err)
		return nil, err
	}
	return role, nil
}

func (r *accountRoleRepositoryImpl) Update(ctx context.Context, role *entity.AccountRole) (*entity.AccountRole, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Model(&entity.AccountRole{}).Where("id = ?", role.Id).Update(transaction, role)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *accountRoleRepositoryImpl) Delete(ctx context.Context, id int64) (err error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	rowsAffected, err := r.baseRepository.Model(&entity.AccountRole{}).Where("id = ?", id).Delete(transaction)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return err
	}
	if rowsAffected == 0 {
		return errors.New(errUtil.ERR_NO_RECORD)
	}
	return nil
}
