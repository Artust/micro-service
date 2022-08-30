package corporation

import (
	pb "avatar/services/gateway/protos/corporation"
	"context"
)

type GetCorporationInput struct {
	Id int64 `json:"id" binding:"required"`
}

func GetCorporation(input *GetCorporationInput, corporationClient pb.CorporationClient) (*CreateCorporationOutPut, error) {
	data := &pb.GetCorporationRequest{
		Id: input.Id,
	}
	response, err := corporationClient.GetCorporation(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &CreateCorporationOutPut{
		Id:        response.Id,
		Name:      response.Name,
		Address:   response.Address,
		Detail:    response.Detail,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	}
	return output, nil
}
