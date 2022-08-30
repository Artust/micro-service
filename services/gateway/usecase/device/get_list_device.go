package device

import (
	pb "avatar/services/gateway/protos/corporation"
	"context"
)

type GetListDeviceInput struct {
	AccountId int64 `form:"accountId"`
	PosId     int64 `form:"posId"`
	CenterId  int64 `form:"centerId"`
	Page      int64 `form:"page"`
	PerPage   int64 `form:"perPage"`
}

type GetListDeviceOutPut struct {
	Results []*GetDeviceOutPut `json:"results"`
}

func GetListDevice(input *GetListDeviceInput, corporationClient pb.CorporationClient) (*GetListDeviceOutPut, error) {
	data := &pb.GetListDeviceRequest{
		Page:      input.Page,
		PerPage:   input.PerPage,
		AccountId: input.AccountId,
		PosId:     input.PosId,
		CenterId:  input.CenterId,
	}
	response, err := corporationClient.GetListDevice(context.Background(), data)
	if err != nil {
		return nil, err
	}
	var output GetListDeviceOutPut
	output.Results = make([]*GetDeviceOutPut, 0)
	for _, val := range response.GetListDeviceResponse {
		output.Results = append(output.Results, &GetDeviceOutPut{
			Id:           val.Id,
			Maker:        val.Maker,
			SerialNumber: val.SerialNumber,
			DeviceType:   val.DeviceType,
			UsePurpose:   val.UsePurpose,
			Owner:        val.Owner,
			OnsiteType:   val.OnsiteType,
			AccountId:    val.AccountId,
			PosId:        val.PosId,
			CenterId:     val.CenterId,
			CreatedAt:    val.CreatedAt,
			UpdatedAt:    val.UpdatedAt,
		})
	}
	return &output, nil
}
