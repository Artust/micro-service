package ip_camera

import (
	corporation "avatar/services/gateway/protos/corporation"
	pos "avatar/services/gateway/protos/pos"
	"context"
)

type GetIpCameraInput struct {
	Id int64 `json:"id" binding:"required"`
}

func GetIpCamera(
	input *GetIpCameraInput,
	posClient pos.POSClient,
	corporationClient corporation.CorporationClient,
) (*CreateIpCameraOutput, error) {
	ipCameraRequest := &pos.GetByIdRequest{
		Id: input.Id,
	}
	ipCameraResponse, err := posClient.GetIpCamera(context.Background(), ipCameraRequest)
	if err != nil {
		return nil, err
	}
	deviceRequest := &corporation.GetDeviceRequest{
		Id: ipCameraResponse.DeviceId,
	}
	deviceResponse, err := corporationClient.GetDevice(context.Background(), deviceRequest)
	if err != nil {
		return nil, err
	}
	return &CreateIpCameraOutput{
		ID:        ipCameraResponse.Id,
		DeviceId:        deviceResponse.Id,
		IsPrimaryCamera: ipCameraResponse.IsPrimaryCamera,
		Maker:           deviceResponse.Maker,
		SerialNumber:    deviceResponse.SerialNumber,
		Resolution:      deviceResponse.Resolution,
		PosId:           deviceResponse.PosId,
		CreatedAt:       ipCameraResponse.CreatedAt,
		PublicURI:       ipCameraResponse.PublicURI,
		PrivateURI:      ipCameraResponse.PrivateURI,
		UpdatedAt:       ipCameraResponse.UpdatedAt,
	}, nil
}
