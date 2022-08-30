package service_template_category

import (
	"context"
	pb "avatar/services/gateway/protos/center"
)

type DeleteCategoryInput struct {
	Id int64
}
type DeleteCategoryOutput struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func DeleteServiceTemplateCategory(input *DeleteCategoryInput, centerClient pb.CenterClient) (*DeleteCategoryOutput, error) {
	ctx := context.Background()
	DeleteCategoryInput := &pb.DeleteByIdRequest{
		Id: input.Id,
	}
	response, err := centerClient.DeleteServiceTemplateCategory(ctx, DeleteCategoryInput)
	if err != nil {
		return nil, err
	}
	output := &DeleteCategoryOutput{
		RowsAffected: response.RowsAffected,
	}

	return output, nil
}
