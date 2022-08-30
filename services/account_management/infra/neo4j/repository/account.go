package repository

import (
	"avatar/pkg/base_repository"
	"avatar/services/account_management/config"
	"avatar/services/account_management/domain/entity"
	"avatar/services/account_management/domain/repository"
	"context"

	pb "avatar/services/account_management/protos"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

type accountRepositoryImpl struct {
	baseRepository base_repository.BaseRepository
}

func NewAccountRepository(
	baseRepository base_repository.BaseRepository,
) repository.AccountRepository {
	return &accountRepositoryImpl{
		baseRepository: baseRepository,
	}
}

func (r *accountRepositoryImpl) GetByEmail(ctx context.Context, email string) (*entity.Account, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	var result entity.Account
	err := r.baseRepository.Model(&entity.Account{}).Where("email = ?", email).Where("status = ?", 0).FindOne(transaction, &result)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return &result, nil
}

func (r *accountRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.Account, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	var result entity.Account
	err := r.baseRepository.Model(&entity.Account{}).Where("id = ?", id).FindOne(transaction, &result)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return &result, nil
}

func (r *accountRepositoryImpl) GetList(ctx context.Context, input *pb.GetListAccountRequest) ([]*entity.Account, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	result := []*entity.Account{}
	query := r.baseRepository.
		Model(&entity.Account{})
	if input.Gender > 0 {
		query = query.Where("gender = ?", input.Gender)
	}
	if input.RoleId > 0 {
		query = query.Where("roleId = ?", input.RoleId)
	}
	if input.CenterId > 0 {
		query = query.Where("centerId = ?", input.CenterId)
	}
	if input.Status > 0 {
		query = query.Where("status = ?", input.Status)
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

func (r *accountRepositoryImpl) Create(ctx context.Context, input *entity.Account) (*entity.Account, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Create(transaction, input)
	if err != nil {
		log.Error("error when write transaction create account, error: ", err)
		return nil, err
	}
	input.Password = ""
	return input, nil
}

func (r *accountRepositoryImpl) Update(ctx context.Context, input *entity.Account) (*entity.Account, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Model(&entity.Account{}).Where("id = ?", input.Id).Update(transaction, input)
	if err != nil {
		return nil, err
	}
	input.Password = ""
	return input, nil
}

func (r *accountRepositoryImpl) ChangePassword(ctx context.Context, input *entity.ChangePasswordData) (err error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	result := entity.Account{
		Password: input.NewPassword,
	}
	err = r.baseRepository.Model(&entity.Account{}).Where("id = ?", input.Id).Update(transaction, &result)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return err
	}
	return nil
}

func (r *accountRepositoryImpl) Deactive(ctx context.Context, id int64) (err error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err = r.baseRepository.Model(&entity.Account{}).Where("id = ?", id).Update(transaction, &entity.Account{Status: 1})
	return err
}

func (r *accountRepositoryImpl) Active(ctx context.Context, id int64) (err error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err = r.baseRepository.Model(&entity.Account{}).Where("id = ?", id).Update(transaction, &entity.Account{Status: 0})
	return err
}
