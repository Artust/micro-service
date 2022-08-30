package device

import (
	pb "avatar/services/gateway/protos/corporation"
	"context"
)

type CreateDeviceInput struct {
	Maker        string `json:"maker" binding:"required"`
	SerialNumber string `json:"serialNumber" binding:"required"`
	DeviceType   string `json:"deviceType" binding:"required"`
	UsePurpose   string `json:"usePurpose" binding:"required"`
	Owner        int64  `json:"owner" binding:"required"`
	OnsiteType   bool   `json:"onsiteType" binding:"required"`
	AccountId    int64  `json:"accountId" binding:"required"`
	PosId        int64  `json:"posId" binding:"required"`
	CenterId     int64  `json:"centerId" binding:"required"`
}

type CreateDeviceOutPut struct {
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
}

func CreateDevice(input *CreateDeviceInput, corporationClient pb.CorporationClient) (*CreateDeviceOutPut, error) {
	data := &pb.CreateDeviceRequest{
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
	response, err := corporationClient.CreateDevice(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &CreateDeviceOutPut{
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
	}
	return output, nil
}
