package rtsp_stream

import (
	"avatar/services/gateway/config"
	"avatar/services/gateway/domain/entity"
	corporation "avatar/services/gateway/protos/corporation"
	pos "avatar/services/gateway/protos/pos"
	"context"
)

type GetRtspStreamOutput struct {
	Ids              []int64 `json:"ids"`
	DefaultsCameraId int64   `json:"defaultsCameraId"`
	PosId            int64   `json:"posId"`
	ServerUri        string  `json:"serverUri"`
}

func GetRtspStream(
	posClient pos.POSClient,
	corporationClient corporation.CorporationClient,
	cfg *config.Environment,
) (*GetRtspStreamOutput, error) {
	ctx := context.Background()
	deviceRequest := corporation.GetListDeviceRequest{
		Page:       1,
		PerPage:    20,
		PosId:      1,
		DeviceType: entity.IpCamera,
	}
	devicesResponse, err := corporationClient.GetListDevice(ctx, &deviceRequest)
	if err != nil {
		return nil, err
	}
	var deviceIds []int64
	for _, device := range devicesResponse.GetListDeviceResponse {
		deviceIds = append(deviceIds, device.Id)
	}
	getListIpCameraRequest := &pos.GetListIpCameraRequest{
		Page:      1,
		PerPage:   20,
		DeviceIds: deviceIds,
	}
	ipCameraResponse, err := posClient.GetListIpCamera(ctx, getListIpCameraRequest)
	if err != nil {
		return nil, err
	}
	var output GetRtspStreamOutput
	for _, val := range ipCameraResponse.GetListIpCameraResponse {
		if val.IsPrimaryCamera {
			output.DefaultsCameraId = val.Id
		}
		output.Ids = append(output.Ids, val.Id)
	}
	output.PosId = 1
	output.ServerUri = cfg.ServerRtspUri
	return &output, nil
}
