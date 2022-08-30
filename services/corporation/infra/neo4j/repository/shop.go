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

type shopRepositoryImpl struct {
	baseRepository base_repository.BaseRepository
}

func NewShopRepository(
	baseRepository base_repository.BaseRepository,
) repository.ShopRepository {
	return &shopRepositoryImpl{
		baseRepository: baseRepository,
	}
}

func (r *shopRepositoryImpl) Create(ctx context.Context, input *entity.Shop) (*entity.Shop, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Create(transaction, input)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *shopRepositoryImpl) GetList(ctx context.Context, input entity.GetListShopOption) ([]*entity.Shop, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	result := []*entity.Shop{}
	query := r.baseRepository.
		Model(&entity.Shop{})
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

func (r *shopRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.Shop, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	var result entity.Shop
	err := r.baseRepository.Model(&entity.Shop{}).Where("id = ?", id).FindOne(transaction, &result)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return &result, nil
}

func (r *shopRepositoryImpl) Update(ctx context.Context, id int64, input *entity.Shop) (*entity.Shop, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Model(&entity.Shop{}).Where("id = ?", id).Update(transaction, input)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *shopRepositoryImpl) Delete(ctx context.Context, id int64) (rowsAffected int64, err error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	rowsAffected, err = r.baseRepository.Model(&entity.Shop{}).Where("id = ?", id).Delete(transaction)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return 0, err
	}
	return rowsAffected, nil
}
