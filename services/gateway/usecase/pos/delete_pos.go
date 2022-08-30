package pos

import (
	pb "avatar/services/gateway/protos/pos"
	"context"
)

type DeletePosInput struct {
	Id int64 `json:"id"`
}

type DeletePosOutPut struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func DeletePos(input *DeletePosInput, posClient pb.POSClient) (*DeletePosOutPut, error) {
	data := &pb.DeleteByIdRequest{
		Id: input.Id,
	}
	response, err := posClient.DeletePos(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &DeletePosOutPut{
		RowsAffected: response.RowsAffected,
	}
	return output, nil
}
