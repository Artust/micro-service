package routine_pos

import (
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/pos"
	"context"
	"fmt"
)

type GetRoutineInput struct {
	Id int64 `json:"id"`
}

func GetRoutine(
	input *GetRoutineInput,
	posClient pb.POSClient,
	cfg *config.Environment) (*CreateRoutineOutput, error) {
	data := &pb.GetByIdRequest{
		Id: int64(input.Id),
	}
	response, err := posClient.GetRoutine(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &CreateRoutineOutput{
		Id:                       response.Id,
		Name:                     response.Name,
		Detail:                   response.Detail,
		AnimationFile:            fmt.Sprintf("%s%s", cfg.S3Uri, response.AnimationFile),
		ImageFile:                fmt.Sprintf("%s%s", cfg.S3Uri, response.ImageFile),
		SoundFile:                fmt.Sprintf("%s%s", cfg.S3Uri, response.SoundFile),
		StartDate:                response.StartDate,
		EndDate:                  response.EndDate,
		PosId:                    response.PosId,
		ServiceTemplateId:        response.ServiceTemplateId,
		CategoryId:               response.CategoryId,
		CreatedAt:                response.CreatedAt,
		UpdatedAt:                response.UpdatedAt,
		ServiceTemplateRoutineId: response.ServiceTemplateRoutineId,
		Gender:                   response.Gender,
	}
	return output, nil
}
