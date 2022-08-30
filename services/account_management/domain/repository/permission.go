package repository

import (
	"avatar/services/account_management/domain/entity"
	"context"
)

type PermissionRepository interface {
	GetById(ctx context.Context, id int64) (*entity.Permission, error)
	GetList(ctx context.Context) ([]*entity.Permission, error)
	Create(ctx context.Context, permission *entity.Permission) (*entity.Permission, error)
	Update(ctx context.Context, permission *entity.Permission) (*entity.Permission, error)
	Delete(ctx context.Context, id int64) error
}
