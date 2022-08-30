package service_template

import (
	pb "avatar/services/gateway/protos/center"
	"context"
)

type DeleteServiceTemplateInput struct {
	Id int64 `json:"id" binding:"required"`
}
type DeleteServiceTemplateOutput struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func DeleteServiceTemplate(input *DeleteServiceTemplateInput, centerClient pb.CenterClient) (*DeleteServiceTemplateOutput, error) {
	ctx := context.Background()
	deleteServiceTemplateInput := &pb.DeleteByIdRequest{
		Id: input.Id,
	}
	response, err := centerClient.DeleteServiceTemplate(ctx, deleteServiceTemplateInput)
	if err != nil {
		return nil, err
	}
	createServiceTemplateOutput := &DeleteServiceTemplateOutput{
		RowsAffected: response.RowsAffected,
	}
	return createServiceTemplateOutput, nil
}
