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

type monitorRepositoryImpl struct {
	baseRepository base_repository.BaseRepository
}

func NewMonitorRepository(
	baseRepository base_repository.BaseRepository,
) repository.MonitorRepository {
	return &monitorRepositoryImpl{
		baseRepository: baseRepository,
	}
}

func (r *monitorRepositoryImpl) Create(ctx context.Context, input *entity.Monitor) (*entity.Monitor, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Create(transaction, input)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *monitorRepositoryImpl) GetList(ctx context.Context, input entity.GetListMonitorOption) ([]*entity.Monitor, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	result := []*entity.Monitor{}
	query := r.baseRepository.
		Model(&entity.Monitor{})
	if input.PosId > 0 {
		query = query.Where("posId = ?", input.PosId)
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

func (r *monitorRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.Monitor, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	var result entity.Monitor
	err := r.baseRepository.Model(&entity.Monitor{}).Where("id = ?", id).FindOne(transaction, &result)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return &result, nil
}

func (r *monitorRepositoryImpl) Update(ctx context.Context, id int64, input *entity.Monitor) (*entity.Monitor, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Model(&entity.Monitor{}).Where("id = ?", id).Select("*").Update(transaction, input)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *monitorRepositoryImpl) Delete(ctx context.Context, id int64) (rowsAffected int64, err error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	rowsAffected, err = r.baseRepository.Model(&entity.Monitor{}).Where("id = ?", id).Delete(transaction)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return 0, err
	}
	return rowsAffected, nil
}
