package device

import (
	"context"
	pb "avatar/services/gateway/protos/corporation"
)

type DeleteDeviceInput struct {
	Id int64 `json:"id" binding:"required"`
}

type DeleteDeviceOutPut struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func DeleteDevice(input *DeleteDeviceInput, corporationClient pb.CorporationClient) (*DeleteDeviceOutPut, error) {
	data := &pb.DeleteDeviceRequest{
		Id: input.Id,
	}
	response, err := corporationClient.DeleteDevice(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &DeleteDeviceOutPut{
		RowsAffected: response.RowsAffected,
	}
	return output, nil
}
