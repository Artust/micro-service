package repository

import (
	"avatar/pkg/base_repository"
	"avatar/services/corporation/config"
	"avatar/services/corporation/domain/entity"
	"avatar/services/corporation/domain/repository"
	"context"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

type centerRepositoryImpl struct {
	baseRepository base_repository.BaseRepository
}

func NewcenterRepository(
	baseRepository base_repository.BaseRepository,
) repository.CenterRepository {
	return &centerRepositoryImpl{
		baseRepository: baseRepository,
	}
}

func (r *centerRepositoryImpl) Create(ctx context.Context, input *entity.Center) (*entity.Center, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Create(transaction, input)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *centerRepositoryImpl) GetList(ctx context.Context, input entity.GetListCenterOption) ([]*entity.Center, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	result := []*entity.Center{}
	query := r.baseRepository.Model(&entity.Center{})
	if input.PerPage > 0 {
		query = query.Limit(input.PerPage)
	}
	if input.Page > 0 {
		query = query.Skip((input.Page - 1) * input.PerPage)
	}
	if input.CorporationId > 0 {
		query = query.Where("corporationId = ?", input.CorporationId)
	}
	err := query.Find(transaction, &result)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return result, nil
}

func (r *centerRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.Center, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	var result entity.Center
	err := r.baseRepository.Model(&entity.Center{}).Where("id = ?", id).FindOne(transaction, &result)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return &result, nil
}

func (r *centerRepositoryImpl) Update(ctx context.Context, id int64, input *entity.Center) (*entity.Center, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Model(&entity.Center{}).Where("id = ?", id).Update(transaction, input)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *centerRepositoryImpl) Delete(ctx context.Context, id int64) (rowsAffected int64, err error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	rowsAffected, err = r.baseRepository.Model(&entity.Center{}).Where("id = ?", id).Delete(transaction)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return 0, err
	}
	return rowsAffected, nil
}
