package monitor

import (
	pb "avatar/services/gateway/protos/pos"
	"context"
)

type DeleteMonitorInput struct {
	Id int64 `json:"id"`
}

type DeleteMonitorOutput struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func DeleteMonitor(input *DeleteMonitorInput,  posClient pb.POSClient) (*DeleteMonitorOutput, error) {
	data := &pb.DeleteByIdRequest{
		Id: input.Id,
	}
	res, err := posClient.DeleteMonitor(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &DeleteMonitorOutput{
		RowsAffected: res.RowsAffected,
	}
	return output, nil
}
