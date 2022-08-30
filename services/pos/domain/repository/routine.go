package repository

import (
	"avatar/services/pos/domain/entity"
	"context"
)

type RoutineRepository interface {
	Create(ctx context.Context, input *entity.PosRoutine) (*entity.PosRoutine, error)
	GetList(ctx context.Context, query entity.GetListRoutineOption) ([]*entity.PosRoutine, error)
	GetListRoutineCategory(ctx context.Context, query entity.GetListRoutineByCategoryOption) ([]*entity.PosRoutine, error)
	GetById(ctx context.Context, id int64) (*entity.PosRoutine, error)
	Update(ctx context.Context, id int64, input *entity.PosRoutine) (*entity.PosRoutine, error)
	Delete(ctx context.Context, option *entity.DeleteRoutineOption) (rowsAffected int64, err error)
	CreateMany(ctx context.Context, input *[]entity.PosRoutine) (*[]entity.PosRoutine, error)
}
