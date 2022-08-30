package shop

import (
	pb "avatar/services/gateway/protos/corporation"
	"context"
)

type GetListShopInput struct {
	Page    int64 `form:"page"`
	PerPage int64 `form:"perPage"`
}

type GetListShopOutPut struct {
	Results []*CreateShopOutPut `json:"results"`
}

func GetListShop(input *GetListShopInput, corporationClient pb.CorporationClient) (*GetListShopOutPut, error) {
	data := &pb.GetListShopRequest{
		Page:    input.Page,
		PerPage: input.PerPage,
	}
	response, err := corporationClient.GetListShop(context.Background(), data)
	if err != nil {
		return nil, err
	}
	var output GetListShopOutPut
	output.Results = make([]*CreateShopOutPut, 0)
	for _, val := range response.GetListShopResponse {
		output.Results = append(output.Results, &CreateShopOutPut{
			Id:        val.Id,
			Name:      val.Name,
			Address:   val.Address,
			CreatedBy: val.CreatedBy,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}
	return &output, nil
}
