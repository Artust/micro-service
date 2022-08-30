package repository

import (
	"avatar/services/pos/domain/entity"
	"context"
)

type PosRepository interface {
	Create(ctx context.Context, input *entity.Pos) (*entity.Pos, error)
	GetList(ctx context.Context, query entity.GetListPosOption) ([]*entity.Pos, error)
	GetById(ctx context.Context, id int64) (*entity.Pos, error)
	Update(ctx context.Context, id int64, input *entity.Pos) (*entity.Pos, error)
	Delete(ctx context.Context, id int64) (rowsAffected int64, err error)
}
