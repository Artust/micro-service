package repository

import (
	"avatar/services/center/domain/entity"
	"context"
)

type ServiceTemplateRepository interface {
	Create(ctx context.Context, input *entity.ServiceTemplate) (*entity.ServiceTemplate, error)
	GetList(ctx context.Context, query entity.GetListServiceTemplateOption) ([]*entity.ServiceTemplate, error)
	GetById(ctx context.Context, id int64) (*entity.ServiceTemplate, error)
	Update(ctx context.Context, id int64, input *entity.ServiceTemplate) (*entity.ServiceTemplate, error)
	Delete(ctx context.Context, id int64) (rowsAffected int64, err error)
}
