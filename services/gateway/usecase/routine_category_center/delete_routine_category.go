package routine_category_center

import (
	"context"
	pb "avatar/services/gateway/protos/center"
)

type DeleteRoutineCategoryInput struct {
	Id int64
}
type DeleteRoutineCategoryOutput struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func DeleteRoutineCategory(input *DeleteRoutineCategoryInput, centerClient pb.CenterClient) (*DeleteRoutineCategoryOutput, error) {
	ctx := context.Background()
	deleteRoutineCategoryInput := &pb.DeleteByIdRequest{
		Id: input.Id,
	}
	response, err := centerClient.DeleteRoutineCategory(ctx, deleteRoutineCategoryInput)
	if err != nil {
		return nil, err
	}
	deleteRoutineCategoryOutput := &DeleteRoutineCategoryOutput{
		RowsAffected: response.RowsAffected,
	}

	return deleteRoutineCategoryOutput, nil
}
