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

type deviceRepositoryImpl struct {
	baseRepository base_repository.BaseRepository
}

func NewDeviceRepository(
	baseRepository base_repository.BaseRepository,
) repository.DeviceRepository {
	return &deviceRepositoryImpl{
		baseRepository: baseRepository,
	}
}

func (r *deviceRepositoryImpl) Create(ctx context.Context, input *entity.Device) (*entity.Device, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Create(transaction, input)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *deviceRepositoryImpl) GetList(ctx context.Context, input entity.GetListDeviceOption) ([]*entity.Device, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	result := []*entity.Device{}
	query := r.baseRepository.
		Model(&entity.Device{})
	if input.AccountId > 0 {
		query = query.Where("accountId = ?", input.AccountId)
	}
	if input.CenterId > 0 {
		query = query.Where("centerId = ?", input.CenterId)
	}
	if input.PosId > 0 {
		query = query.Where("posId = ?", input.PosId)
	}
	if input.DeviceType != "" {
		query = query.Where("deviceType = ?", input.DeviceType)
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

func (r *deviceRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.Device, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	var result entity.Device
	err := r.baseRepository.Model(&entity.Device{}).Where("id = ?", id).FindOne(transaction, &result)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return &result, nil
}

func (r *deviceRepositoryImpl) Update(ctx context.Context, id int64, input *entity.Device) (*entity.Device, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Model(&entity.Device{}).Where("id = ?", id).Update(transaction, input)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *deviceRepositoryImpl) Delete(ctx context.Context, id int64) (rowsAffected int64, err error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	rowsAffected, err = r.baseRepository.Model(&entity.Device{}).Where("id = ?", id).Delete(transaction)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return 0, err
	}
	return rowsAffected, nil
}
