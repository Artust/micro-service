package center

import (
	pb "avatar/services/gateway/protos/corporation"
	"context"
)

type UpdateCenterInput struct {
	Id            int64
	Name          string `json:"name"`
	Detail        string `json:"detail"`
	Type          string `json:"type"`
	CorporationId int64  `json:"corporationId"`
}

type UpdateCenterOutPut struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	Detail        string `json:"detail"`
	Type          string `json:"type"`
	CorporationId int64  `json:"corporationId"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}

func UpdateCenter(input *UpdateCenterInput, corporationClient pb.CorporationClient) (*UpdateCenterOutPut, error) {
	data := &pb.UpdateCenterRequest{
		Id:            input.Id,
		Name:          input.Name,
		Detail:        input.Detail,
		Type:          input.Type,
		CorporationId: input.CorporationId,
	}
	response, err := corporationClient.UpdateCenter(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &UpdateCenterOutPut{
		Id:            response.Id,
		Name:          response.Name,
		Detail:        response.Detail,
		Type:          response.Type,
		CorporationId: response.CorporationId,
		CreatedAt:     response.CreatedAt,
		UpdatedAt:     response.UpdatedAt,
	}
	return output, nil
}
