package center

import (
	pb "avatar/services/gateway/protos/corporation"
	"context"
)

type GetCenterInput struct {
	ID int64
}

type GetCenterOutPut struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	Detail        string `json:"detail"`
	Type          string `json:"type"`
	CorporationId int64  `json:"corporationId"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}

func GetCenter(input *GetCenterInput, corporationClient pb.CorporationClient) (*GetCenterOutPut, error) {
	data := &pb.GetCenterRequest{
		Id: input.ID,
	}
	response, err := corporationClient.GetCenter(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &GetCenterOutPut{
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
