package corporation

import (
	pb "avatar/services/gateway/protos/corporation"
	"context"
)

type GetListCorporationInput struct {
	Page    int64 `form:"page"`
	PerPage int64 `form:"perPage"`
}

type GetListCorporationOutPut struct {
	Results []CreateCorporationOutPut `json:"results"`
}

func GetListCorporation(input *GetListCorporationInput, corporationClient pb.CorporationClient) (*GetListCorporationOutPut, error) {
	data := &pb.GetListCorporationRequest{
		Page:    input.Page,
		PerPage: input.PerPage,
	}
	response, err := corporationClient.GetListCorporation(context.Background(), data)
	if err != nil {
		return nil, err
	}
	var output GetListCorporationOutPut
	for _, val := range response.GetListCorporationResponse {
		output.Results = append(output.Results, CreateCorporationOutPut{
			Id:        val.Id,
			Name:      val.Name,
			Address:   val.Address,
			Detail:    val.Detail,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}
	return &output, nil
}
