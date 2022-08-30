package shop

import (
	pb "avatar/services/gateway/protos/corporation"
	"context"
)

type GetShopInput struct {
	Id int64 `json:"id"`
}

func GetShop(input *GetShopInput, corporationClient pb.CorporationClient) (*CreateShopOutPut, error) {
	data := &pb.GetShopRequest{
		Id: input.Id,
	}
	response, err := corporationClient.GetShop(context.Background(), data)
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
