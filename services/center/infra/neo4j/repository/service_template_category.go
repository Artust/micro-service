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

type serviceTemplateCategoryRepository struct {
	baseRepository base_repository.BaseRepository
}

func NewServiceTemplateCategoryRepository(
	baseRepository base_repository.BaseRepository,
) repository.ServiceTemplateCategoryRepository {
	return &serviceTemplateCategoryRepository{
		baseRepository: baseRepository,
	}
}

func (r *serviceTemplateCategoryRepository) Create(ctx context.Context, input *entity.ServiceTemplateCategory) (*entity.ServiceTemplateCategory, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Create(transaction, input)
	if err != nil {
		log.Error("error when create service template category, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *serviceTemplateCategoryRepository) GetList(ctx context.Context, input entity.GetListServiceTemplateCategoryOption) ([]*entity.ServiceTemplateCategory, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	result := []*entity.ServiceTemplateCategory{}
	query := r.baseRepository.Model(&entity.ServiceTemplateCategory{})
	if input.PerPage > 0 {
		query = query.Limit(input.PerPage)
	}
	if input.Page > 0 {
		query = query.Skip((input.Page - 1) * input.PerPage)
	}
	err := query.Find(transaction, &result)
	if err != nil {
		log.Error("error when get list service template category, error: ", err)
		return nil, err
	}
	return result, nil
}

func (r *serviceTemplateCategoryRepository) GetById(ctx context.Context, id int64) (*entity.ServiceTemplateCategory, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	var result entity.ServiceTemplateCategory
	err := r.baseRepository.Model(&entity.ServiceTemplateCategory{}).Where("id = ?", id).FindOne(transaction, &result)
	if err != nil {
		log.Error("error when get service template category, error: ", err)
		return nil, err
	}
	return &result, nil
}

func (r *serviceTemplateCategoryRepository) Update(ctx context.Context, id int64, input *entity.ServiceTemplateCategory) (*entity.ServiceTemplateCategory, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Model(&entity.ServiceTemplateCategory{}).Where("id = ?", id).Update(transaction, input)
	if err != nil {
		log.Error("error when update service template category, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *serviceTemplateCategoryRepository) Delete(ctx context.Context, id int64) (rowsAffected int64, err error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	rowsAffected, err = r.baseRepository.Model(&entity.ServiceTemplateCategory{}).Where("id = ?", id).Delete(transaction)
	if err != nil {
		log.Error("error when delete service template category, error: ", err)
		return 0, err
	}
	return rowsAffected, nil
}
