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

type routineRepositoryImpl struct {
	baseRepository base_repository.BaseRepository
}

func NewroutineRepository(
	baseRepository base_repository.BaseRepository,
) repository.RoutineRepository {
	return &routineRepositoryImpl{
		baseRepository: baseRepository,
	}
}

func (r *routineRepositoryImpl) GetListRoutineCategory(ctx context.Context, input entity.GetListRoutineOption) ([]*entity.CenterRoutine, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	results := []*entity.CenterRoutine{}
	query := r.baseRepository.
		Model(&entity.CenterRoutine{})
	if len(input.CategoryIds) > 0 {
		query = query.Where("categoryId IN ?", input.CategoryIds)
	}
	if len(input.Ids) > 0 {
		query = query.Where("id IN ?", input.Ids)
	}
	if input.PerPage > 0 {
		query = query.Limit(input.PerPage)
	}
	if input.Page > 0 {
		query = query.Skip((input.Page - 1) * input.PerPage)
	}
	err := query.Find(transaction, &results)
	if err != nil {
		log.Error("error when get list routine by category, error: ", err)
		return nil, err
	}
	return results, nil
}

func (r *routineRepositoryImpl) Create(ctx context.Context, input *entity.CenterRoutine) (*entity.CenterRoutine, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Create(transaction, input)
	if err != nil {
		log.Error("error when create routine, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *routineRepositoryImpl) GetList(ctx context.Context, input entity.GetListRoutineOption) ([]*entity.CenterRoutine, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	result := []*entity.CenterRoutine{}
	query := r.baseRepository.Model(&entity.CenterRoutine{})
	if input.PerPage > 0 {
		query = query.Limit(input.PerPage)
	}
	if input.Page > 0 {
		query = query.Skip((input.Page - 1) * input.PerPage)
	}
	if input.CategoryId > 0 {
		query = query.Where("categoryId = ?", input.CategoryId)
	}
	if input.Gender > 0 {
		query = query.Where("gender = ?", input.Gender)
	}
	if len(input.Ids) > 0 {
		query = query.Where("id IN ?", input.Ids)
	}
	query.Order("createdAt DESC")
	err := query.Find(transaction, &result)
	if err != nil {
		log.Error("error when get list routine, error: ", err)
		return nil, err
	}
	return result, nil
}

func (r *routineRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.CenterRoutine, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	var result entity.CenterRoutine
	err := r.baseRepository.Model(&entity.CenterRoutine{}).Where("id = ?", id).FindOne(transaction, &result)
	if err != nil {
		log.Error("error when get routine, error: ", err)
		return nil, err
	}
	return &result, nil
}

func (r *routineRepositoryImpl) Update(ctx context.Context, id int64, input *entity.CenterRoutine) (*entity.CenterRoutine, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Model(&entity.CenterRoutine{}).Where("id = ?", id).Select("*").Update(transaction, input)
	if err != nil {
		log.Error("error when update routine, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *routineRepositoryImpl) Delete(ctx context.Context, id int64) (rowsAffected int64, err error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	rowsAffected, err = r.baseRepository.Model(&entity.CenterRoutine{}).Where("id = ?", id).Delete(transaction)
	if err != nil {
		log.Error("error when delete routine, error: ", err)
		return 0, err
	}
	return rowsAffected, nil
}
