package routine_category_pos

import (
	pb "avatar/services/gateway/protos/pos"
	"context"
)

type GetRoutineCategoryInput struct {
	Id int64 `json:"id"`
}

func GetRoutineCategory(input *GetRoutineCategoryInput, posClient pb.POSClient) (*CreateRoutineCategoryOutput, error) {
	data := &pb.GetByIdRequest{
		Id: input.Id,
	}
	response, err := posClient.GetRoutineCategory(context.Background(), data)
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
