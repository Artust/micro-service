package service_template_category

import (
	pb "avatar/services/gateway/protos/center"
	"context"
)

type CreateCategoryInput struct {
	Name string `json:"name" binding:"required"`
}
type CreateCategoryOutput struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func CreateServiceTemplateCategory(input *CreateCategoryInput, centerClient pb.CenterClient) (*CreateCategoryOutput, error) {
	ctx := context.Background()
	createServiceTemplateCategoryInput := &pb.CreateServiceTemplateCategoryRequest{
		Name: input.Name,
	}
	response, err := centerClient.CreateServiceTemplateCategory(ctx, createServiceTemplateCategoryInput)
	if err != nil {
		return nil, err
	}
	output := &CreateCategoryOutput{
		Id:        response.Id,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	}

	return output, nil
}
