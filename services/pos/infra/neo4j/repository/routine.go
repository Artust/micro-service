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

type routineRepositoryImpl struct {
	baseRepository base_repository.BaseRepository
}

func NewRoutineRepository(
	baseRepository base_repository.BaseRepository,
) repository.RoutineRepository {
	return &routineRepositoryImpl{
		baseRepository: baseRepository,
	}
}

func (r *routineRepositoryImpl) Create(ctx context.Context, input *entity.PosRoutine) (*entity.PosRoutine, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Create(transaction, input)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *routineRepositoryImpl) CreateMany(ctx context.Context, input *[]entity.PosRoutine) (*[]entity.PosRoutine, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.CreateMany(transaction, input)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *routineRepositoryImpl) GetList(ctx context.Context, input entity.GetListRoutineOption) ([]*entity.PosRoutine, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	result := []*entity.PosRoutine{}
	query := r.baseRepository.
		Model(&entity.PosRoutine{})
	if input.PosId > 0 {
		query = query.Where("posId = ?", input.PosId)
	}
	if input.Gender > 0 {
		query = query.Where("gender = ?", input.Gender)
	}
	if len(input.Ids) > 0 {
		query = query.Where("id in ", input.Ids)
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

func (r *routineRepositoryImpl) GetListRoutineCategory(ctx context.Context, input entity.GetListRoutineByCategoryOption) ([]*entity.PosRoutine, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	results := []*entity.PosRoutine{}
	query := r.baseRepository.
		Model(&entity.PosRoutine{})
	if input.EndDate.Year() > 1 {
		if input.Between == 0 {
			query = query.Where("endDate > ?", input.EndDate)
		} else {
			query = query.Where("endDate < ?", input.EndDate)
		}
	}
	if input.PosId > 0 {
		query = query.Where("posId = ?", input.PosId)
	}
	if len(input.ListCategoryId) > 0 {
		query = query.Where("categoryId in ", input.ListCategoryId)
	}
	if len(input.Ids) > 0 {
		query = query.Where("id in ", input.Ids)
	}
	if input.PerPage > 0 {
		query = query.Limit(input.PerPage)
	}
	if input.Page > 0 {
		query = query.Skip((input.Page - 1) * input.PerPage)
	}
	err := query.Find(transaction, &results)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return results, nil
}

func (r *routineRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.PosRoutine, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	var result entity.PosRoutine
	err := r.baseRepository.Model(&entity.PosRoutine{}).Where("id = ?", id).FindOne(transaction, &result)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return &result, nil
}

func (r *routineRepositoryImpl) Update(ctx context.Context, id int64, input *entity.PosRoutine) (*entity.PosRoutine, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Model(&entity.PosRoutine{}).Where("id = ?", id).Select("*").Update(transaction, input)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *routineRepositoryImpl) Delete(ctx context.Context, option *entity.DeleteRoutineOption) (rowsAffected int64, err error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	query := r.baseRepository.Model(&entity.PosRoutine{})
	if option.Id > 0 {
		query.Where("id = ?", option.Id)
	}
	if option.PosId > 0 {
		query.Where("posId = ?", option.PosId)
	}
	rowsAffected, err = query.Delete(transaction)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return 0, err
	}
	return rowsAffected, nil
}
