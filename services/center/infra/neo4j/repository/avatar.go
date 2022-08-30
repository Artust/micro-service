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

type avatarRepository struct {
	baseRepository base_repository.BaseRepository
}

func NewAvatarRepository(
	baseRepository base_repository.BaseRepository,
) repository.AvatarRepository {
	return &avatarRepository{
		baseRepository: baseRepository,
	}
}

func (r *avatarRepository) Create(ctx context.Context, input *entity.Avatar) (*entity.Avatar, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Create(transaction, input)
	if err != nil {
		log.Error("error when create avatar, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *avatarRepository) GetList(ctx context.Context, input entity.GetListAvatarOption) ([]*entity.Avatar, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	result := []*entity.Avatar{}
	query := r.baseRepository.Model(&entity.Avatar{})
	if input.PerPage > 0 {
		query = query.Limit(input.PerPage)
	}
	if input.Page > 0 {
		query = query.Skip((input.Page - 1) * input.PerPage)
	}
	if input.Gender > 0 {
		query = query.Where("gender = ?", input.Gender)
	}
	if len(input.Ids) > 0 {
		query = query.Where("id IN ?", input.Ids)
	}
	err := query.Find(transaction, &result)
	if err != nil {
		log.Error("error when get list avatar, error: ", err)
		return nil, err
	}
	return result, nil
}

func (r *avatarRepository) GetById(ctx context.Context, id int64) (*entity.Avatar, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	var result entity.Avatar
	err := r.baseRepository.Model(&entity.Avatar{}).Where("id = ?", id).FindOne(transaction, &result)
	if err != nil {
		log.Error("error when get avatar, error: ", err)
		return nil, err
	}
	return &result, nil
}

func (r *avatarRepository) Update(ctx context.Context, id int64, input *entity.Avatar) (*entity.Avatar, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Model(&entity.Avatar{}).Where("id = ?", id).Select("*").Update(transaction, input)
	if err != nil {
		log.Error("error when update avatar, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *avatarRepository) Delete(ctx context.Context, id int64) (rowsAffected int64, err error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	rowsAffected, err = r.baseRepository.Model(&entity.Avatar{}).Where("id = ?", id).Delete(transaction)
	if err != nil {
		log.Error("error when delete avatar, error: ", err)
		return 0, err
	}
	return rowsAffected, nil
}
