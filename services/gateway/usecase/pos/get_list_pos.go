package pos

import (
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/pos"
	"context"
	"fmt"
)

type GetListPosInput struct {
	ShopId   int64 `form:"shopId"`
	CenterId int64 `form:"centerId"`
	Page     int64 `form:"page"`
	PerPage  int64 `form:"perPage"`
}

type GetListPosOutPut struct {
	Results []*CreatePosOutput `json:"result"`
}

func GetListPos(
	input *GetListPosInput,
	posClient pb.POSClient,
	cfg *config.Environment,
) (*GetListPosOutPut, error) {
	data := &pb.GetListPosRequest{
		ShopId:   input.ShopId,
		CenterId: input.CenterId,
		Page:     input.Page,
		PerPage:  input.PerPage,
	}
	response, err := posClient.GetListPos(context.Background(), data)
	if err != nil {
		return nil, err
	}
	var output GetListPosOutPut
	output.Results = make([]*CreatePosOutput, 0)
	for _, Pos := range response.GetListPosResponse {
		backgroundName := []string{}
		for _, background := range Pos.Backgrounds {
			backgroundName = append(backgroundName, fmt.Sprintf("%s/%s/%s", cfg.S3Uri, config.BackgroundBucketName, background))
		}
		output.Results = append(output.Results, &CreatePosOutput{
			Id:                      Pos.Id,
			Name:                    Pos.Name,
			ServiceName:             Pos.ServiceName,
			ServiceCategory:         Pos.ServiceCategory,
			ServiceDetail:           Pos.ServiceDetail,
			ShopId:                  Pos.ShopId,
			CenterId:                Pos.CenterId,
			ServiceTemplateId:       Pos.ServiceTemplateId,
			MaleRoutineIds:          Pos.MaleRoutineIds,
			DefaultMaleRoutineIds:   Pos.DefaultMaleRoutineIds,
			FemaleRoutineIds:        Pos.FemaleRoutineIds,
			DefaultFemaleRoutineIds: Pos.DefaultFemaleRoutineIds,
			DefaultAvatarId:     Pos.DefaultAvatarId,
			UseServiceTemplate:      Pos.UseServiceTemplate,
			Backgrounds:             backgroundName,
			DefaultBackground:       fmt.Sprintf("%s/%s/%s", cfg.S3Uri, config.BackgroundBucketName, Pos.DefaultBackground),
			CreatedBy:               Pos.CreatedBy,
			CreatedAt:               Pos.CreatedAt,
			UpdatedAt:               Pos.UpdatedAt,
		})
	}
	return &output, nil
}
