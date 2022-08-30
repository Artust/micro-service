package monitor

import (
	pb "avatar/services/gateway/protos/pos"
	"context"
)

type CreateMonitorInput struct {
	Maker              string `json:"maker" binding:"required"`
	SerialNumber       string `json:"serialNumber" binding:"required"`
	MonitorStatus      string `json:"monitorStatus"`
	ResolutionWidth    int64  `json:"resolutionWidth" binding:"required"`
	ResolutionHeight   int64  `json:"resolutionHeight" binding:"required"`
	HorizontalRotation bool   `json:"horizontalRotation"`
	PosId              int64  `json:"posId" binding:"required"`
}

type CreateMonitorOutput struct {
	Id                 int64  `json:"id"`
	Maker              string `json:"maker"`
	SerialNumber       string `json:"serialNumber"`
	MonitorStatus      string `json:"monitorStatus"`
	ResolutionWidth    int64  `json:"resolutionWidth"`
	ResolutionHeight   int64  `json:"resolutionHeight"`
	HorizontalRotation bool   `json:"horizontalRotation"`
	PosId              int64  `json:"posId"`
	CreatedAt          string `json:"createdAt"`
	UpdatedAt          string `json:"updatedAt,omitempty"`
}

func CreateMonitor(input *CreateMonitorInput, posClient pb.POSClient) (*CreateMonitorOutput, error) {
	data := &pb.CreateMonitorRequest{
		Maker:              input.Maker,
		SerialNumber:       input.SerialNumber,
		MonitorStatus:      input.MonitorStatus,
		ResolutionWidth:    input.ResolutionWidth,
		ResolutionHeight:   input.ResolutionHeight,
		HorizontalRotation: input.HorizontalRotation,
		PosId:              input.PosId,
	}
	response, err := posClient.CreateMonitor(context.Background(), data)
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
	}
	return output, nil
}
