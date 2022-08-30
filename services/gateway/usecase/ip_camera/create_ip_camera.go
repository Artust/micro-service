package ip_camera

import (
	"avatar/services/gateway/domain/entity"
	corporation "avatar/services/gateway/protos/corporation"
	pos "avatar/services/gateway/protos/pos"
	"context"
)

type CreateIpCameraInput struct {
	IsPrimaryCamera bool   `json:"isPrimaryCamera"`
	Maker           string `json:"maker" binding:"required"`
	SerialNumber    string `json:"serialNumber" binding:"required"`
	Resolution      string `json:"resolution" binding:"required"`
	PublicURI       string `json:"publicURI" binding:"required"`
	PrivateURI      string `json:"privateURI" binding:"required"`
}

type CreateIpCameraOutput struct {
	ID              int64  `json:"id"`
	DeviceId        int64  `json:"deviceId"`
	IsPrimaryCamera bool   `json:"isPrimaryCamera"`
	Maker           string `json:"maker"`
	SerialNumber    string `json:"serialNumber"`
	Resolution      string `json:"resolution"`
	PosId           int64  `json:"posId"`
	CreatedAt       string `json:"createdAt"`
	PublicURI       string `json:"publicURI"`
	PrivateURI      string `json:"privateURI"`
	UpdatedAt       string `json:"updatedAt,omitempty"`
}

func CreateIpCamera(
	input *CreateIpCameraInput,
	posClient pos.POSClient,
	corporationClient corporation.CorporationClient,
) (*CreateIpCameraOutput, error) {
	ctx := context.Background()
	createDeviceRequest := &corporation.CreateDeviceRequest{
		Maker:        input.Maker,
		SerialNumber: input.SerialNumber,
		DeviceType:   entity.IpCamera,
		UsePurpose:   "",
		Owner:        0,
		OnsiteType:   true,
		AccountId:    0,
		PosId:        1,
		CenterId:     0,
		Resolution:   input.Resolution,
	}
	device, err := corporationClient.CreateDevice(ctx, createDeviceRequest)
	if err != nil {
		return nil, err
	}
	createIpCameraRequest := &pos.CreateIpCameraRequest{
		IsPrimaryCamera: input.IsPrimaryCamera,
		PublicURI:       input.PublicURI,
		PrivateURI:      input.PrivateURI,
		DeviceId:        device.Id,
	}
	ipCamera, err := posClient.CreateIpCamera(ctx, createIpCameraRequest)
	if err != nil {
		return nil, err
	}
	output := &CreateIpCameraOutput{
		ID:              ipCamera.Id,
		DeviceId:        device.Id,
		IsPrimaryCamera: ipCamera.IsPrimaryCamera,
		Maker:           device.Maker,
		SerialNumber:    device.SerialNumber,
		Resolution:      device.Resolution,
		PosId:           device.PosId,
		CreatedAt:       ipCamera.CreatedAt,
		PublicURI:       ipCamera.PublicURI,
		PrivateURI:      ipCamera.PrivateURI,
		UpdatedAt:       ipCamera.CreatedAt,
	}
	return output, nil
}
