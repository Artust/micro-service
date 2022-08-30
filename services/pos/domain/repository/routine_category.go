package repository

import (
	"avatar/services/pos/domain/entity"
	"context"
)

type RoutineCategoryRepository interface {
	Create(ctx context.Context, input *entity.RoutineCategory) (*entity.RoutineCategory, error)
	GetList(ctx context.Context, query entity.GetListRoutineCategoryOption) ([]*entity.RoutineCategory, error)
	GetById(ctx context.Context, id int64) (*entity.RoutineCategory, error)
	Update(ctx context.Context, id int64, input *entity.RoutineCategory) (*entity.RoutineCategory, error)
	Delete(ctx context.Context, id int64) (rowsAffected int64, err error)
}
