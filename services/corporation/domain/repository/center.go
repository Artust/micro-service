package repository

import (
	"avatar/services/corporation/domain/entity"
	"context"
)

type CenterRepository interface {
	Create(ctx context.Context, input *entity.Center) (*entity.Center, error)
	GetList(ctx context.Context, query entity.GetListCenterOption) ([]*entity.Center, error)
	GetById(ctx context.Context, id int64) (*entity.Center, error)
	Update(ctx context.Context, id int64, input *entity.Center) (*entity.Center, error)
	Delete(ctx context.Context, id int64) (rowsAffected int64, err error)
}
