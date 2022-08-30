package center

import (
	pb "avatar/services/gateway/protos/corporation"
	"context"
)

type CreateCenterInput struct {
	Name          string `json:"name" binding:"required"`
	Detail        string `json:"detail" binding:"required"`
	Type          string `json:"type" binding:"required"`
	CorporationId int64  `json:"corporationId" binding:"required"`
}

type CreateCenterOutPut struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	Detail        string `json:"detail"`
	Type          string `json:"type"`
	CorporationId int64  `json:"corporationId"`
	CreatedAt     string `json:"createdAt"`
}

func CreateCenter(input *CreateCenterInput, corporationClient pb.CorporationClient) (*CreateCenterOutPut, error) {
	data := &pb.CreateCenterRequest{
		Name:          input.Name,
		Detail:        input.Detail,
		Type:          input.Type,
		CorporationId: input.CorporationId,
	}
	response, err := corporationClient.CreateCenter(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &CreateCenterOutPut{
		Id:            response.Id,
		Name:          response.Name,
		Detail:        response.Detail,
		Type:          response.Type,
		CorporationId: response.CorporationId,
		CreatedAt:     response.CreatedAt,
	}
	return output, nil
}
