package ip_camera

import (
	corporation "avatar/services/gateway/protos/corporation"
	pos "avatar/services/gateway/protos/pos"
	"context"
)

type DeleteIpCameraInput struct {
	Id int64 `json:"id"`
}

type DeleteIpCameraOutput struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func DeleteIpCamera(
	input *DeleteIpCameraInput,
	posClient pos.POSClient,
	corporationClient corporation.CorporationClient,
) (*DeleteIpCameraOutput, error) {
	ipCameraId := &pos.DeleteByIdRequest{
		Id: input.Id,
	}
	response, err := posClient.DeleteIpCamera(context.Background(), ipCameraId)
	if err != nil {
		return nil, err
	}
	return &DeleteIpCameraOutput{
		RowsAffected: response.RowsAffected,
	}, nil
}
