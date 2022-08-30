package service_template_category

import (
	pb "avatar/services/gateway/protos/center"
	"context"
)

type GetCategoryInput struct {
	Id int64
}
type GetCategoryOutput struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func GetServiceTemplateCategory(input *GetCategoryInput, centerClient pb.CenterClient) (*GetCategoryOutput, error) {
	ctx := context.Background()
	GetCategoryInput := &pb.GetByIdRequest{
		Id: input.Id,
	}
	response, err := centerClient.GetServiceTemplateCategory(ctx, GetCategoryInput)
	if err != nil {
		return nil, err
	}
	output := &GetCategoryOutput{
		Id:        response.Id,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	}

	return output, nil
}
