package shop

import (
	pb "avatar/services/gateway/protos/corporation"
	"context"
)

type CreateShopInput struct {
	Name      string `json:"name" binding:"required"`
	Address   string `json:"address" binding:"required"`
	CreatedBy int64
}

type CreateShopOutPut struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	CreatedBy int64  `json:"createdBy"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

func CreateShop(input *CreateShopInput, corporationClient pb.CorporationClient) (*CreateShopOutPut, error) {
	data := &pb.CreateShopRequest{
		Name:      input.Name,
		Address:   input.Address,
		CreatedBy: input.CreatedBy,
	}
	response, err := corporationClient.CreateShop(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &CreateShopOutPut{
		Id:        response.Id,
		Name:      response.Name,
		Address:   response.Address,
		CreatedBy: response.CreatedBy,
		CreatedAt: response.CreatedAt,
	}
	return output, nil
}
