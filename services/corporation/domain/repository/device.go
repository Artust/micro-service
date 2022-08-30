package repository

import (
	"avatar/services/corporation/domain/entity"
	"context"
)

type DeviceRepository interface {
	Create(ctx context.Context, input *entity.Device) (*entity.Device, error)
	GetList(ctx context.Context, query entity.GetListDeviceOption) ([]*entity.Device, error)
	GetById(ctx context.Context, id int64) (*entity.Device, error)
	Update(ctx context.Context, id int64, input *entity.Device) (*entity.Device, error)
	Delete(ctx context.Context, id int64) (rowsAffected int64, err error)
}
