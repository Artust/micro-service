package repository

import (
	"avatar/services/account_management/domain/entity"
	"context"
)

type AccountRoleRepository interface {
	GetById(ctx context.Context, id int64) (*entity.AccountRole, error)
	GetList(ctx context.Context) ([]*entity.AccountRole, error)
	Create(ctx context.Context, role *entity.AccountRole) (*entity.AccountRole, error)
	Update(ctx context.Context, role *entity.AccountRole) (*entity.AccountRole, error)
	Delete(ctx context.Context, id int64) error
}
