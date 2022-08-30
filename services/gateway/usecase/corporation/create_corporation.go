package corporation

import (
	pb "avatar/services/gateway/protos/corporation"
	"context"
)

type CreateCorporationInput struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Detail  string `json:"detail" binding:"required"`
}

type CreateCorporationOutPut struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Detail    string `json:"detail"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

func CreateCorporation(input *CreateCorporationInput, corporationClient pb.CorporationClient) (*CreateCorporationOutPut, error) {
	data := &pb.CreateCorporationRequest{
		Name:    input.Name,
		Address: input.Address,
		Detail:  input.Detail,
	}
	response, err := corporationClient.CreateCorporation(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &CreateCorporationOutPut{
		Id:        response.Id,
		Name:      response.Name,
		Address:   response.Address,
		Detail:    response.Detail,
		CreatedAt: response.CreatedAt,
	}
	return output, nil
}
