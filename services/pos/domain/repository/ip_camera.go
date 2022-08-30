package repository

import (
	"avatar/services/pos/domain/entity"
	"context"
)

type IpCameraRepository interface {
	Create(ctx context.Context, input *entity.IpCamera) (*entity.IpCamera, error)
	GetList(ctx context.Context, query entity.GetListIpCameraOption) ([]*entity.IpCamera, error)
	GetById(ctx context.Context, id int64) (*entity.IpCamera, error)
	Update(ctx context.Context, id int64, input *entity.IpCamera) (*entity.IpCamera, error)
	Delete(ctx context.Context, id int64) (rowsAffected int64, err error)
}
