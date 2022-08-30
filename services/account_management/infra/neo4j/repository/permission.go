package repository

import (
	"avatar/pkg/base_repository"
	errUtil "avatar/pkg/err"
	"avatar/services/account_management/config"
	"avatar/services/account_management/domain/entity"
	"avatar/services/account_management/domain/repository"
	"context"
	"errors"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

type permissionRepositoryImpl struct {
	baseRepository base_repository.BaseRepository
}

func NewPermissionRepository(
	baseRepository base_repository.BaseRepository,
) repository.PermissionRepository {
	return &permissionRepositoryImpl{
		baseRepository: baseRepository,
	}
}

func (r *permissionRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.Permission, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	var result entity.Permission
	err := r.baseRepository.Model(&entity.Permission{}).Where("id = ?", id).FindOne(transaction, &result)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	fmt.Println("Permission: ", &result)
	return &result, nil
}

func (r *permissionRepositoryImpl) GetList(ctx context.Context) ([]*entity.Permission, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	result := []*entity.Permission{}
	err := r.baseRepository.Model(&entity.Permission{}).Find(transaction, &result)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return result, nil
}

func (r *permissionRepositoryImpl) Create(ctx context.Context, permission *entity.Permission) (*entity.Permission, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Create(transaction, permission)
	if err != nil {
		log.Error("error when write transaction create permission, error: ", err)
		return nil, err
	}
	return permission, nil
}

func (r *permissionRepositoryImpl) Update(ctx context.Context, permission *entity.Permission) (*entity.Permission, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Model(&entity.Permission{}).Where("id = ?", permission.Id).Update(transaction, permission)
	if err != nil {
		return nil, err
	}
	return permission, nil
}

func (r *permissionRepositoryImpl) Delete(ctx context.Context, id int64) (err error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	rowsAffected, err := r.baseRepository.Model(&entity.Permission{}).Where("id = ?", id).Delete(transaction)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return err
	}
	if rowsAffected == 0 {
		return errors.New(errUtil.ERR_NO_RECORD)
	}
	return nil
}
