package corporation

import (
	pb "avatar/services/gateway/protos/corporation"
	"context"
)

type DeleteCorporationInput struct {
	Id int64 `json:"id" binding:"required"`
}

type DeleteCorporationOutPut struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func DeleteCorporation(input *DeleteCorporationInput, corporationClient pb.CorporationClient) (*DeleteCorporationOutPut, error) {
	data := &pb.DeleteCorporationRequest{
		Id: input.Id,
	}
	response, err := corporationClient.DeleteCorporation(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &DeleteCorporationOutPut{
		RowsAffected: response.RowsAffected,
	}
	return output, nil
}
