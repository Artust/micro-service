package monitor

import (
	pb "avatar/services/gateway/protos/pos"
	"context"
)

type GetListMonitorInput struct {
	Page     int64 `form:"page"`
	PerPage  int64 `form:"perPage"`
	PosId    int64 `form:"posId"`
}

type GetListMonitorOutput struct {
	Results []CreateMonitorOutput `json:"results"`
}

func GetListMonitor(input *GetListMonitorInput, posClient pb.POSClient) (*GetListMonitorOutput, error) {
	data := &pb.GetListMonitorRequest{
		Page:     input.Page,
		PerPage:  input.PerPage,
		PosId:    input.PosId,
	}
	response, err := posClient.GetListMonitor(context.Background(), data)
	if err != nil {
		return nil, err
	}
	var output GetListMonitorOutput
	for _, val := range response.GetListMonitorResponse {
		output.Results = append(output.Results, CreateMonitorOutput{
			Id:                 val.Id,
			Maker:              val.Maker,
			SerialNumber:       val.SerialNumber,
			MonitorStatus:      val.MonitorStatus,
			ResolutionWidth:    val.ResolutionWidth,
			ResolutionHeight:   val.ResolutionHeight,
			HorizontalRotation: val.HorizontalRotation,
			PosId:              val.PosId,
			CreatedAt:          val.CreatedAt,
			UpdatedAt:          val.UpdatedAt,
		})
	}
	return &output, nil
}
