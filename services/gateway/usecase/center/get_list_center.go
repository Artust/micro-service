package center

import (
	pb "avatar/services/gateway/protos/corporation"
	"context"
)

type GetListCenterInput struct {
	CorporationId int64 `form:"corporationId"`
	Page          int64 `form:"page"`
	PerPage       int64 `form:"perPage"`
}

type GetListCenterOutPut struct {
	Results []*GetCenterOutPut `json:"results"`
}

func GetListCenter(input *GetListCenterInput, corporationClient pb.CorporationClient) (*GetListCenterOutPut, error) {
	data := &pb.GetListCenterRequest{
		Page:          input.Page,
		PerPage:       input.PerPage,
		CorporationId: input.CorporationId,
	}
	response, err := corporationClient.GetListCenter(context.Background(), data)
	if err != nil {
		return nil, err
	}
	var output GetListCenterOutPut
	output.Results = make([]*GetCenterOutPut, 0)
	for _, center := range response.ListCenter {
		output.Results = append(output.Results, &GetCenterOutPut{
			Id:            center.Id,
			Name:          center.Name,
			Detail:        center.Detail,
			Type:          center.Type,
			CorporationId: center.CorporationId,
			CreatedAt:     center.CreatedAt,
			UpdatedAt:     center.UpdatedAt,
		})
	}
	return &output, nil
}
