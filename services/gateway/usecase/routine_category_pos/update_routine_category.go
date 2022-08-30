package routine_category_pos

import (
	pb "avatar/services/gateway/protos/pos"
	"context"
)

type UpdateRoutineCategoryInput struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func UpdateRoutineCategory(input *UpdateRoutineCategoryInput, posClient pb.POSClient) (*CreateRoutineCategoryOutput, error) {
	data := &pb.UpdateRoutineCategoryRequest{
		Id:   input.Id,
		Name: input.Name,
	}
	response, err := posClient.UpdateRoutineCategory(context.Background(), data)
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
