package shop

import (
	pb "avatar/services/gateway/protos/corporation"
	"context"
)

type UpdateShopInput struct {
	Id      int64
	Name    string `json:"name"`
	Address string `json:"address"`
}

func UpdateShop(input *UpdateShopInput, corporationClient pb.CorporationClient) (*CreateShopOutPut, error) {
	data := &pb.UpdateShopRequest{
		Id:       input.Id,
		Name:     input.Name,
		Address:  input.Address,
	}
	response, err := corporationClient.UpdateShop(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &CreateShopOutPut{
		Id:        response.Id,
		Name:      response.Name,
		Address:   response.Address,
		CreatedBy: response.CreatedBy,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	}
	return output, nil
}
