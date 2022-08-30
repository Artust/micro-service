package repository

import (
	"avatar/services/pos/domain/entity"
	"context"
)

type MonitorRepository interface {
	Create(ctx context.Context, input *entity.Monitor) (*entity.Monitor, error)
	GetList(ctx context.Context, query entity.GetListMonitorOption) ([]*entity.Monitor, error)
	GetById(ctx context.Context, id int64) (*entity.Monitor, error)
	Update(ctx context.Context, id int64, input *entity.Monitor) (*entity.Monitor, error)
	Delete(ctx context.Context, id int64) (rowsAffected int64, err error)
}
