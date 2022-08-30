package routine_category_pos

import (
	pb "avatar/services/gateway/protos/pos"
	"context"
)

type CreateRoutineCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

type CreateRoutineCategoryOutput struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

func CreateRoutineCategory(input *CreateRoutineCategoryInput, posClient pb.POSClient) (*CreateRoutineCategoryOutput, error) {
	data := &pb.CreateRoutineCategoryRequest{
		Name: input.Name,
	}
	response, err := posClient.CreateRoutineCategory(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &CreateRoutineCategoryOutput{
		Id:        response.Id,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	}
	return output, nil
}
