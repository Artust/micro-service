package repository

import (
	"avatar/services/talk_session/domain/entity"
	"context"
)

type NoteRepository interface {
	Create(ctx context.Context, input *entity.Note) (*entity.Note, error)
	GetList(ctx context.Context, query entity.GetListNoteOption) ([]*entity.Note, error)
	GetById(ctx context.Context, id int64) (*entity.Note, error)
	Update(ctx context.Context, id int64, input *entity.Note) (*entity.Note, error)
	Delete(ctx context.Context, id int64) (rowsAffected int64, err error)
}
