package ip_camera

import (
	"avatar/services/gateway/domain/entity"
	corporation "avatar/services/gateway/protos/corporation"
	pos "avatar/services/gateway/protos/pos"
	"context"
)

type UpdateIpCameraInput struct {
	CameraId        int64  `json:"cameraId"`
	DeviceId        int64  `json:"deviceId"`
	IsPrimaryCamera bool   `json:"isPrimaryCamera"`
	Maker           string `json:"maker"`
	Resolution      string `json:"resolution"`
	PublicURI       string `json:"publicURI"`
	PrivateURI      string `json:"privateURI"`
	SerialNumber    string `json:"serialNumber"`
}

func UpdateIpCamera(
	input *UpdateIpCameraInput,
	posClient pos.POSClient,
	corporationClient corporation.CorporationClient,
) (*CreateIpCameraOutput, error) {
	ctx := context.Background()
	updateIpCameraRequest := &pos.UpdateIpCameraRequest{
		Id:              input.CameraId,
		IsPrimaryCamera: input.IsPrimaryCamera,
		PublicURI:       input.PublicURI,
		PrivateURI:      input.PrivateURI,
		DeviceId:        input.DeviceId,
	}
	ipCameraResponse, err := posClient.UpdateIpCamera(ctx, updateIpCameraRequest)
	if err != nil {
		return nil, err
	}
	updateDeviceRequest := &corporation.UpdateDeviceRequest{
		Id:           input.DeviceId,
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
	deviceResponse, err := corporationClient.UpdateDevice(ctx, updateDeviceRequest)
	if err != nil {
		return nil, err
	}
	return &CreateIpCameraOutput{
		ID:              ipCameraResponse.Id,
		DeviceId:        ipCameraResponse.DeviceId,
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
