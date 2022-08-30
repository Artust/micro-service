package shop

import (
	"context"
	pb "avatar/services/gateway/protos/corporation"
)

type DeleteShopInput struct {
	Id int64
}

type DeleteShopOutPut struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func DeleteShop(input *DeleteShopInput, corporationClient pb.CorporationClient) (*DeleteShopOutPut, error) {
	data := &pb.DeleteShopRequest{
		Id: input.Id,
	}
	response, err := corporationClient.DeleteShop(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &DeleteShopOutPut{
		RowsAffected: response.RowsAffected,
	}
	return output, nil
}
