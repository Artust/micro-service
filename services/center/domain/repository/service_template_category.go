package repository

import (
	"avatar/services/center/domain/entity"
	"context"
)

type ServiceTemplateCategoryRepository interface {
	Create(ctx context.Context, input *entity.ServiceTemplateCategory) (*entity.ServiceTemplateCategory, error)
	GetList(ctx context.Context, query entity.GetListServiceTemplateCategoryOption) ([]*entity.ServiceTemplateCategory, error)
	GetById(ctx context.Context, id int64) (*entity.ServiceTemplateCategory, error)
	Update(ctx context.Context, id int64, input *entity.ServiceTemplateCategory) (*entity.ServiceTemplateCategory, error)
	Delete(ctx context.Context, id int64) (rowsAffected int64, err error)
}
