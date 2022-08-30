package repository

import (
	"avatar/services/account_management/domain/entity"
	"context"
)

type UserActivityRepository interface {
	Get(ctx context.Context, accountId int64) ([]*entity.UserActivity, error)
	Save(ctx context.Context, activity *entity.UserActivity) (*entity.UserActivity, error)
}
