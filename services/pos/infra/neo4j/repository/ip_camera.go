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

type ipCameraRepositoryImpl struct {
	baseRepository base_repository.BaseRepository
}

func NewIpCameraRepository(
	baseRepository base_repository.BaseRepository,
) repository.IpCameraRepository {
	return &ipCameraRepositoryImpl{
		baseRepository: baseRepository,
	}
}

func (r *ipCameraRepositoryImpl) Create(ctx context.Context, input *entity.IpCamera) (*entity.IpCamera, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Create(transaction, input)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *ipCameraRepositoryImpl) GetList(ctx context.Context, input entity.GetListIpCameraOption) ([]*entity.IpCamera, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	result := []*entity.IpCamera{}
	query := r.baseRepository.
		Model(&entity.IpCamera{})
	if input.PosId > 0 {
		query = query.Where("posId = ?", input.PosId)
	}
	if len(input.DeviceId) > 0 {
		query = query.Where("deviceId IN ?", input.DeviceId)
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

func (r *ipCameraRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.IpCamera, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	var result entity.IpCamera
	err := r.baseRepository.Model(&entity.IpCamera{}).Where("id = ?", id).FindOne(transaction, &result)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return &result, nil
}

func (r *ipCameraRepositoryImpl) Update(ctx context.Context, id int64, input *entity.IpCamera) (*entity.IpCamera, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Model(&entity.IpCamera{}).Where("id = ?", id).Select("*").Update(transaction, input)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *ipCameraRepositoryImpl) Delete(ctx context.Context, id int64) (rowsAffected int64, err error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	rowsAffected, err = r.baseRepository.Model(&entity.IpCamera{}).Where("id = ?", id).Delete(transaction)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return 0, err
	}
	return rowsAffected, nil
}
