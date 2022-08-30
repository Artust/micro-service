package repository

import (
	"avatar/pkg/base_repository"
	"avatar/services/talk_session/config"
	"avatar/services/talk_session/domain/entity"
	"avatar/services/talk_session/domain/repository"
	"context"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

type noteRepositoryImpl struct {
	baseRepository base_repository.BaseRepository
}

func NewNoteRepository(
	baseRepository base_repository.BaseRepository,
) repository.NoteRepository {
	return &noteRepositoryImpl{
		baseRepository: baseRepository,
	}
}

func (r *noteRepositoryImpl) Create(ctx context.Context, input *entity.Note) (*entity.Note, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Create(transaction, input)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *noteRepositoryImpl) GetList(ctx context.Context, input entity.GetListNoteOption) ([]*entity.Note, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	result := []*entity.Note{}
	query := r.baseRepository.
		Model(&entity.Note{})
	if input.TalkSessionId > 0 {
		query = query.Where("talkSessionId = ?", input.TalkSessionId)
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

func (r *noteRepositoryImpl) Update(ctx context.Context, id int64, input *entity.Note) (*entity.Note, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Model(&entity.Note{}).Where("id = ?", id).Select("*").Update(transaction, input)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *noteRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.Note, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	var result entity.Note
	err := r.baseRepository.Model(&entity.Note{}).Where("id = ?", id).FindOne(transaction, &result)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	return &result, nil
}

func (r *noteRepositoryImpl) Delete(ctx context.Context, id int64) (rowsAffected int64, err error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	rowsAffected, err = r.baseRepository.Model(&entity.Note{}).Where("id = ?", id).Delete(transaction)
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return 0, err
	}
	return rowsAffected, nil
}
