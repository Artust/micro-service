package device

import (
	pb "avatar/services/gateway/protos/corporation"
	"context"
)

type UpdateDeviceInput struct {
	Id           int64  `json:"id"`
	Maker        string `json:"maker"`
	SerialNumber string `json:"serialNumber"`
	DeviceType   string `json:"deviceType"`
	UsePurpose   string `json:"usePurpose"`
	Owner        int64  `json:"owner"`
	OnsiteType   bool   `json:"onsiteType"`
	AccountId    int64  `json:"accountId"`
	PosId        int64  `json:"posId"`
	CenterId     int64  `json:"centerId"`
}

type UpdateDeviceOutPut struct {
	Id           int64  `json:"id"`
	Maker        string `json:"maker"`
	SerialNumber string `json:"serialNumber"`
	DeviceType   string `json:"deviceType"`
	UsePurpose   string `json:"usePurpose"`
	Owner        int64  `json:"owner"`
	OnsiteType   bool   `json:"onsiteType"`
	AccountId    int64  `json:"accountId"`
	PosId        int64  `json:"posId"`
	CenterId     int64  `json:"centerId"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

func UpdateDevice(input *UpdateDeviceInput, corporationClient pb.CorporationClient) (*UpdateDeviceOutPut, error) {
	data := &pb.UpdateDeviceRequest{
		Id:           input.Id,
		Maker:        input.Maker,
		SerialNumber: input.SerialNumber,
		DeviceType:   input.DeviceType,
		UsePurpose:   input.UsePurpose,
		Owner:        input.Owner,
		OnsiteType:   input.OnsiteType,
		AccountId:    input.AccountId,
		PosId:        input.PosId,
		CenterId:     input.CenterId,
	}
	response, err := corporationClient.UpdateDevice(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &UpdateDeviceOutPut{
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
