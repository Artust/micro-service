package repository

import (
	"avatar/services/corporation/domain/entity"
	"context"
)

type CorporationRepository interface {
	Create(ctx context.Context, input *entity.Corporation) (*entity.Corporation, error)
	GetList(ctx context.Context, query entity.GetListCorporationOption) ([]*entity.Corporation, error)
	GetById(ctx context.Context, id int64) (*entity.Corporation, error)
	Update(ctx context.Context, id int64, input *entity.Corporation) (*entity.Corporation, error)
	Delete(ctx context.Context, id int64) (rowsAffected int64, err error)
}
