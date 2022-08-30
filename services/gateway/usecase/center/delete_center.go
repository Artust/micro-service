package center

import (
	pb "avatar/services/gateway/protos/corporation"
	"context"
)

type DeleteCenterInput struct {
	ID int64
}

type DeleteCenterOutPut struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func DeleteCenter(input *DeleteCenterInput, corporationClient pb.CorporationClient) (*DeleteCenterOutPut, error) {
	data := &pb.DeleteCenterRequest{
		Id: input.ID,
	}
	response, err := corporationClient.DeleteCenter(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &DeleteCenterOutPut{
		RowsAffected: response.RowsAffected,
	}
	return output, nil
}
