package routine_pos

import (
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/pos"
	"context"
	"fmt"
)

type GetListRoutineInput struct {
	PosId      int64 `form:"posId"`
	CategoryId int64 `form:"categoryId"`
	Page       int64 `form:"page"`
	PerPage    int64 `form:"perPage"`
	Gender     int64 `form:"gender" binding:"min=0,max=1"`
}
type GetListRoutineOutPut struct {
	Results []*CreateRoutineOutput `json:"results"`
}

func GetListRoutine(
	input *GetListRoutineInput,
	posClient pb.POSClient,
	cfg *config.Environment,
) (*GetListRoutineOutPut, error) {
	data := &pb.GetListRoutineRequest{
		PosId:      input.PosId,
		CategoryId: input.CategoryId,
		Page:       input.Page,
		PerPage:    input.PerPage,
		Gender:     input.Gender,
	}
	response, err := posClient.GetListRoutine(context.Background(), data)
	if err != nil {
		return nil, err
	}
	var output GetListRoutineOutPut
	output.Results = make([]*CreateRoutineOutput, 0)
	for _, val := range response.GetListRoutineResponse {
		output.Results = append(output.Results, &CreateRoutineOutput{
			Id:                       val.Id,
			Name:                     val.Name,
			Detail:                   val.Detail,
			AnimationFile:            fmt.Sprintf("%s%s", cfg.S3Uri, val.AnimationFile),
			ImageFile:                fmt.Sprintf("%s%s", cfg.S3Uri, val.ImageFile),
			SoundFile:                fmt.Sprintf("%s%s", cfg.S3Uri, val.SoundFile),
			StartDate:                val.StartDate,
			EndDate:                  val.EndDate,
			PosId:                    val.PosId,
			ServiceTemplateId:        val.ServiceTemplateId,
			CategoryId:               val.CategoryId,
			CreatedAt:                val.CreatedAt,
			UpdatedAt:                val.UpdatedAt,
			ServiceTemplateRoutineId: val.ServiceTemplateRoutineId,
			Gender:                   val.Gender,
		})
	}
	return &output, nil
}
