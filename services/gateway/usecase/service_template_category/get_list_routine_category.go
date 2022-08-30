package service_template_category

import (
	pb "avatar/services/gateway/protos/center"
	"context"
)

type GetListCategoryInput struct {
	Page    int64 `form:"page"`
	PerPage int64 `form:"perPage"`
}
type GetListCategoryOutput struct {
	ListServiceTemplateCategory []*ServiceTemplateCategory `json:"results"`
}
type ServiceTemplateCategory struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func GetListServiceTemplateCategory(input *GetListCategoryInput, centerClient pb.CenterClient) (*GetListCategoryOutput, error) {
	ctx := context.Background()

	GetListCategoryInput := &pb.GetListServiceTemplateCategoryRequest{
		Page:    int64(input.Page),
		PerPage: int64(input.PerPage),
	}
	response, err := centerClient.GetListServiceTemplateCategory(ctx, GetListCategoryInput)
	if err != nil {
		return nil, err
	}
	output := &GetListCategoryOutput{}
	output.ListServiceTemplateCategory = make([]*ServiceTemplateCategory, 0)
	for _, v := range response.GetListServiceTemplateCategoryResponse {
		ServiceTemplateCategory := &ServiceTemplateCategory{
			Id:        v.Id,
			Name:      v.Name,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		output.ListServiceTemplateCategory = append(output.ListServiceTemplateCategory, ServiceTemplateCategory)
		ServiceTemplateCategory = nil
	}
	return output, nil
}
