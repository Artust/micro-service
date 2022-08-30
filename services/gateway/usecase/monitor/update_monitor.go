package monitor

import (
	pb "avatar/services/gateway/protos/pos"
	"context"
)

type UpdateMonitorInput struct {
	Id                 int64  `json:"id" binding:"required"`
	Maker              string `json:"maker" binding:"required"`
	SerialNumber       string `json:"serialNumber" binding:"required"`
	MonitorStatus      string `json:"monitorStatus"`
	ResolutionWidth    int64  `json:"resolutionWidth" binding:"required"`
	ResolutionHeight   int64  `json:"resolutionHeight" binding:"required"`
	HorizontalRotation bool   `json:"horizontalRotation"`
}

func UpdateMonitor(input *UpdateMonitorInput, posClient pb.POSClient) (*CreateMonitorOutput, error) {
	data := &pb.UpdateMonitorRequest{
		Id:                 input.Id,
		Maker:              input.Maker,
		SerialNumber:       input.SerialNumber,
		MonitorStatus:      input.MonitorStatus,
		ResolutionWidth:    input.ResolutionWidth,
		ResolutionHeight:   input.ResolutionHeight,
		HorizontalRotation: input.HorizontalRotation,
	}
	response, err := posClient.UpdateMonitor(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &CreateMonitorOutput{
		Id:                 response.Id,
		Maker:              response.Maker,
		SerialNumber:       response.SerialNumber,
		MonitorStatus:      response.MonitorStatus,
		ResolutionWidth:    response.ResolutionWidth,
		ResolutionHeight:   response.ResolutionHeight,
		HorizontalRotation: response.HorizontalRotation,
		PosId:              response.PosId,
		CreatedAt:          response.CreatedAt,
		UpdatedAt:          response.UpdatedAt,
	}
	return output, nil
}
