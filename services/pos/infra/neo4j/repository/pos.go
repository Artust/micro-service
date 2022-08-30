package repository

import (
	"avatar/pkg/base_repository"
	"avatar/services/pos/config"
	"avatar/services/pos/domain/entity"
	"avatar/services/pos/domain/repository"
	"context"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

type posRepositoryImpl struct {
	baseRepository base_repository.BaseRepository
}

func NewPosRepository(
	baseRepository base_repository.BaseRepository,
) repository.PosRepository {
	return &posRepositoryImpl{
		baseRepository: baseRepository,
	}
}

func (r *posRepositoryImpl) Create(ctx context.Context, input *entity.Pos) (*entity.Pos, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Create(transaction, input)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *posRepositoryImpl) GetList(ctx context.Context, input entity.GetListPosOption) ([]*entity.Pos, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	result := []*entity.Pos{}
	query := r.baseRepository.
		Model(&entity.Pos{})
	if input.CenterId > 0 {
		query = query.Where("centerId = ?", input.CenterId)
	}
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

func (r *posRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.Pos, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	var result entity.Pos
	err := r.baseRepository.Model(&entity.Pos{}).Where("id = ?", id).FindOne(transaction, &result)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return &result, nil
}

func (r *posRepositoryImpl) Update(ctx context.Context, id int64, input *entity.Pos) (*entity.Pos, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Model(&entity.Pos{}).Where("id = ?", id).Select("*").Update(transaction, input)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *posRepositoryImpl) Delete(ctx context.Context, id int64) (rowsAffected int64, err error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	rowsAffected, err = r.baseRepository.Model(&entity.Pos{}).Where("id = ?", id).Delete(transaction)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return 0, err
	}
	return rowsAffected, nil
}
