package repository

import (
	"avatar/pkg/base_repository"
	"avatar/services/center/config"
	"avatar/services/center/domain/entity"
	"avatar/services/center/domain/repository"
	"context"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

type routineCategoryRepository struct {
	baseRepository base_repository.BaseRepository
}

func NewRoutineCategoryRepository(
	baseRepository base_repository.BaseRepository,
) repository.RoutineCategoryRepository {
	return &routineCategoryRepository{
		baseRepository: baseRepository,
	}
}

func (r *routineCategoryRepository) Create(ctx context.Context, input *entity.RoutineCategory) (*entity.RoutineCategory, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Create(transaction, input)
	if err != nil {
		log.Error("error when create routine category, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *routineCategoryRepository) GetList(ctx context.Context, input entity.GetListRoutineCategoryOption) ([]*entity.RoutineCategory, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	result := []*entity.RoutineCategory{}
	query := r.baseRepository.Model(&entity.RoutineCategory{})
	if input.PerPage > 0 {
		query = query.Limit(input.PerPage)
	}
	if input.Page > 0 {
		query = query.Skip((input.Page - 1) * input.PerPage)
	}
	err := query.Find(transaction, &result)
	if err != nil {
		log.Error("error when get list routine category, error: ", err)
		return nil, err
	}
	return result, nil
}

func (r *routineCategoryRepository) GetById(ctx context.Context, id int64) (*entity.RoutineCategory, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	var result entity.RoutineCategory
	err := r.baseRepository.Model(&entity.RoutineCategory{}).Where("id = ?", id).FindOne(transaction, &result)
	if err != nil {
		log.Error("error when get routine category, error: ", err)
		return nil, err
	}
	return &result, nil
}

func (r *routineCategoryRepository) Update(ctx context.Context, id int64, input *entity.RoutineCategory) (*entity.RoutineCategory, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Model(&entity.RoutineCategory{}).Where("id = ?", id).Update(transaction, input)
	if err != nil {
		log.Error("error when update routine category, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *routineCategoryRepository) Delete(ctx context.Context, id int64) (rowsAffected int64, err error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	rowsAffected, err = r.baseRepository.Model(&entity.RoutineCategory{}).Where("id = ?", id).Delete(transaction)
	if err != nil {
		log.Error("error when delete routine category, error: ", err)
		return 0, err
	}
	return rowsAffected, nil
}
