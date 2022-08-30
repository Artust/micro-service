package device

import (
	"context"
	pb "avatar/services/gateway/protos/corporation"
)

type GetDeviceInput struct {
	Id int64 `json:"id" binding:"required"`
}

type GetDeviceOutPut struct {
	Id           int64  `json:"id"`
	Maker        string `json:"maker"`
	SerialNumber string `json:"serialNumber"`
	DeviceType   string `json:"deviceType"`
	UsePurpose   string `json:"usePurpose"`
	Owner        int64  `json:"owner"`
	OnsiteType   bool `json:"onsiteType"`
	AccountId    int64  `json:"accountId"`
	PosId        int64  `json:"posId"`
	CenterId     int64  `json:"centerId"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt,omitempty"`
}

func GetDevice(input *GetDeviceInput, corporationClient pb.CorporationClient) (*GetDeviceOutPut, error) {
	data := &pb.GetDeviceRequest{
		Id: input.Id,
	}
	response, err := corporationClient.GetDevice(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &GetDeviceOutPut{
		Id:           response.Id,
		Maker:        response.Maker,
		SerialNumber: response.SerialNumber,
		DeviceType:   response.DeviceType,
		UsePurpose:   response.UsePurpose,
		Owner:        response.Owner,
		OnsiteType:   response.OnsiteType,
		AccountId:    response.AccountId,
		PosId:        response.PosId,
		CenterId:     response.CenterId,
		CreatedAt:    response.CreatedAt,
		UpdatedAt:    response.UpdatedAt,
	}
	return output, nil
}
