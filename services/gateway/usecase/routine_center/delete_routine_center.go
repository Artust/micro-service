package routine_center

import (
	pb "avatar/services/gateway/protos/center"
	"context"
)

type DeleteRoutineCenterInput struct {
	ID uint64
}
type DeleteRoutineCenterOutput struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func DeleteRoutineCenter(input *DeleteRoutineCenterInput, centerClient pb.CenterClient) (*DeleteRoutineCenterOutput, error) {
	deleteRoutineKeyInput := &pb.DeleteByIdRequest{
		Id: int64(input.ID),
	}
	response, err := centerClient.DeleteRoutine(context.Background(), deleteRoutineKeyInput)
	if err != nil {
		return nil, err
	}

	deleteRoutineCenterOutput := &DeleteRoutineCenterOutput{
		RowsAffected: response.RowsAffected,
	}

	return deleteRoutineCenterOutput, nil
}
