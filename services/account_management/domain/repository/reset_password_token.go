package repository

import (
	"avatar/services/account_management/domain/entity"
	"context"
)

type ResetPasswordTokenRepository interface {
	Create(ctx context.Context, input *entity.ResetPasswordToken) (*entity.ResetPasswordToken, error)
	GetByToken(ctx context.Context, token string) (*entity.ResetPasswordToken, error)
	Delete(ctx context.Context, id int64) (err error)
}
