package monitor

import (
	pb "avatar/services/gateway/protos/pos"
	"context"
)

type GetMonitorInput struct {
	Id int64 `json:"id" binding:"required"`
}

func GetMonitor(input *GetMonitorInput, posClient pb.POSClient) (*CreateMonitorOutput, error) {
	data := &pb.GetByIdRequest{
		Id: input.Id,
	}
	response, err := posClient.GetMonitor(context.Background(), data)
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
