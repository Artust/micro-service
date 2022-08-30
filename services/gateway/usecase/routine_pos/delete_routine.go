package routine_pos

import (
	pb "avatar/services/gateway/protos/pos"
	"context"
)

type DeleteRoutineInput struct {
	Id int64
}
type DeleteRoutineOutPut struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func DeleteRoutine(input *DeleteRoutineInput, posClient pb.POSClient) (*DeleteRoutineOutPut, error) {
	data := &pb.DeleteRoutineRequest{
		Id: input.Id,
	}
	response, err := posClient.DeleteRoutine(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &DeleteRoutineOutPut{
		RowsAffected: response.RowsAffected,
	}
	return output, nil
}
