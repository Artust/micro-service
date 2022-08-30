package ip_camera

import (
	"avatar/services/gateway/domain/entity"
	corporation "avatar/services/gateway/protos/corporation"
	pos "avatar/services/gateway/protos/pos"
	"context"
)

type GetListIpCameraInput struct {
	Page    int64 `form:"page"`
	PerPage int64 `form:"perPage"`
	PosId   int64 `form:"posId"`
}

type GetListIpCameraOutput struct {
	Results []*CreateIpCameraOutput `json:"results"`
}

func GetListIpCamera(
	input *GetListIpCameraInput,
	posClient pos.POSClient,
	corporationClient corporation.CorporationClient,
) (*GetListIpCameraOutput, error) {
	ctx := context.Background()
	getListDeviceRequest := &corporation.GetListDeviceRequest{
		Page:       input.Page,
		PerPage:    input.PerPage,
		PosId:      input.PosId,
		DeviceType: entity.IpCamera,
	}
	deviceResponse, err := corporationClient.GetListDevice(ctx, getListDeviceRequest)
	if err != nil {
		return nil, err
	}
	var deviceIds []int64
	for _, device := range deviceResponse.GetListDeviceResponse {
		deviceIds = append(deviceIds, device.Id)
	}
	getListIpCameraRequest := &pos.GetListIpCameraRequest{
		Page:      input.Page,
		PerPage:   input.PerPage,
		DeviceIds: deviceIds,
	}
	ipCameraResponse, err := posClient.GetListIpCamera(ctx, getListIpCameraRequest)
	if err != nil {
		return nil, err
	}
	var output GetListIpCameraOutput
	output.Results = make([]*CreateIpCameraOutput, 0)
	for _, device := range deviceResponse.GetListDeviceResponse {
		for _, ipCamera := range ipCameraResponse.GetListIpCameraResponse {
			if device.Id == ipCamera.DeviceId {
				output.Results = append(output.Results, &CreateIpCameraOutput{
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
				})
				continue
			}
		}
	}
	return &output, nil
}
