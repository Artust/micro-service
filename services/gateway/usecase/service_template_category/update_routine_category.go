package service_template_category

import (
	pb "avatar/services/gateway/protos/center"
	"context"
)

type UpdateCategoryInput struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type UpdateCategoryOutput struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func UpdateServiceTemplateCategory(input *UpdateCategoryInput, centerClient pb.CenterClient) (*UpdateCategoryOutput, error) {
	ctx := context.Background()

	UpdateCategoryInput := &pb.UpdateServiceTemplateCategoryRequest{
		Id:   input.Id,
		Name: input.Name,
	}
	response, err := centerClient.UpdateServiceTemplateCategory(ctx, UpdateCategoryInput)
	if err != nil {
		return nil, err
	}
	output := &UpdateCategoryOutput{
		Id:        response.Id,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	}

	return output, nil
}
