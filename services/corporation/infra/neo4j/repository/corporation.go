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

type corporationRepositoryImpl struct {
	baseRepository base_repository.BaseRepository
}

func NewCorporationRepository(
	baseRepository base_repository.BaseRepository,
) repository.CorporationRepository {
	return &corporationRepositoryImpl{
		baseRepository: baseRepository,
	}
}

func (r *corporationRepositoryImpl) Create(ctx context.Context, input *entity.Corporation) (*entity.Corporation, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Create(transaction, input)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *corporationRepositoryImpl) GetList(ctx context.Context, input entity.GetListCorporationOption) ([]*entity.Corporation, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	result := []*entity.Corporation{}
	query := r.baseRepository.Model(&entity.Corporation{})
	if input.PerPage > 0 {
		query = query.Limit(input.PerPage)
	}
	if input.Page > 0 {
		query = query.Skip((input.Page - 1) * input.PerPage)
	}
	err := query.Find(transaction, &result)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return result, nil
}

func (r *corporationRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.Corporation, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	var result entity.Corporation
	err := r.baseRepository.Model(&entity.Corporation{}).Where("id = ?", id).FindOne(transaction, &result)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return &result, nil
}

func (r *corporationRepositoryImpl) Update(ctx context.Context, id int64, input *entity.Corporation) (*entity.Corporation, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	log.Println("input1: ", input, id)
	err := r.baseRepository.Model(&entity.Corporation{}).Where("id = ?", id).Update(transaction, input)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *corporationRepositoryImpl) Delete(ctx context.Context, id int64) (rowsAffected int64, err error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	rowsAffected, err = r.baseRepository.Model(&entity.Corporation{}).Where("id = ?", id).Delete(transaction)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return 0, err
	}
	return rowsAffected, nil
}
