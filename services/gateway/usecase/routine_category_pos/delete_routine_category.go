package routine_category_pos

import (
	pb "avatar/services/gateway/protos/pos"
	"context"
)

type DeleteRoutineCategoryInput struct {
	Id int64
}

type DeleteRoutineCategoryOutput struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func DeleteRoutineCategory(input *DeleteRoutineCategoryInput, posClient pb.POSClient) (*DeleteRoutineCategoryOutput, error) {
	data := &pb.DeleteByIdRequest{
		Id: input.Id,
	}
	response, err := posClient.DeleteRoutineCategory(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &DeleteRoutineCategoryOutput{
		RowsAffected: response.RowsAffected,
	}
	return output, nil
}
