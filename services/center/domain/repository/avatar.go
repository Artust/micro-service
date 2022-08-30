package repository

import (
	"avatar/services/center/domain/entity"
	"context"
)

type AvatarRepository interface {
	Create(ctx context.Context, input *entity.Avatar) (*entity.Avatar, error)
	GetList(ctx context.Context, query entity.GetListAvatarOption) ([]*entity.Avatar, error)
	GetById(ctx context.Context, id int64) (*entity.Avatar, error)
	Update(ctx context.Context, id int64, input *entity.Avatar) (*entity.Avatar, error)
	Delete(ctx context.Context, id int64) (rowsAffected int64, err error)
}
