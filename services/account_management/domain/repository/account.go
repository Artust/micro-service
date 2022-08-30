package repository

import (
	"avatar/services/account_management/domain/entity"
	pb "avatar/services/account_management/protos"
	"context"
)

type AccountRepository interface {
	GetByEmail(ctx context.Context, email string) (*entity.Account, error)
	GetById(ctx context.Context, id int64) (*entity.Account, error)
	GetList(ctx context.Context, input *pb.GetListAccountRequest) ([]*entity.Account, error)
	Create(ctx context.Context, account *entity.Account) (*entity.Account, error)
	Update(ctx context.Context, account *entity.Account) (*entity.Account, error)
	ChangePassword(ctx context.Context, input *entity.ChangePasswordData) error
	Deactive(ctx context.Context, id int64) error
	Active(ctx context.Context, id int64) error
}
