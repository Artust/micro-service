package corporation

import (
	pb "avatar/services/gateway/protos/corporation"
	"context"
)

type UpdateCorporationInput struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Detail  string `json:"detail"`
}

func UpdateCorporation(input *UpdateCorporationInput, corporationClient pb.CorporationClient) (*CreateCorporationOutPut, error) {
	data := &pb.UpdateCorporationRequest{
		Id:      input.Id,
		Name:    input.Name,
		Address: input.Address,
		Detail:  input.Detail,
	}
	response, err := corporationClient.UpdateCorporation(context.Background(), data)
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
