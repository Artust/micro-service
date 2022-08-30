package repository

import (
	"avatar/services/center/domain/entity"
	"context"
)

type RoutineRepository interface {
	Create(ctx context.Context, input *entity.CenterRoutine) (*entity.CenterRoutine, error)
	GetList(ctx context.Context, query entity.GetListRoutineOption) ([]*entity.CenterRoutine, error)
	GetById(ctx context.Context, id int64) (*entity.CenterRoutine, error)
	Update(ctx context.Context, id int64, input *entity.CenterRoutine) (*entity.CenterRoutine, error)
	Delete(ctx context.Context, id int64) (rowsAffected int64, err error)
	GetListRoutineCategory(ctx context.Context, query entity.GetListRoutineOption) ([]*entity.CenterRoutine, error)
}
