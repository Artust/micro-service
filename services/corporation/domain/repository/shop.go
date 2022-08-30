package repository

import (
	"avatar/services/corporation/domain/entity"
	"context"
)

type ShopRepository interface {
	Create(ctx context.Context, input *entity.Shop) (*entity.Shop, error)
	GetList(ctx context.Context, query entity.GetListShopOption) ([]*entity.Shop, error)
	GetById(ctx context.Context, id int64) (*entity.Shop, error)
	Update(ctx context.Context, id int64, input *entity.Shop) (*entity.Shop, error)
	Delete(ctx context.Context, id int64) (rowsAffected int64, err error)
}
