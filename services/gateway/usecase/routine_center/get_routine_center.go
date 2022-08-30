package routine_center

import (
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/center"
	"context"
	"fmt"
)

type GetRoutineCenterInput struct {
	ID uint64
}

func GetRoutineCenter(
	input *GetRoutineCenterInput,
	centerClient pb.CenterClient,
	cfg *config.Environment,
) (*CreateRoutineCenterOutput, error) {
	getRoutineKeyInput := &pb.GetByIdRequest{
		Id: int64(input.ID),
	}
	response, err := centerClient.GetRoutine(context.Background(), getRoutineKeyInput)
	if err != nil {
		return nil, err
	}
	routine := &CreateRoutineCenterOutput{
		Id:            response.Id,
		Name:          response.Name,
		Detail:        response.Detail,
		AnimationFile: fmt.Sprintf("%s%s", cfg.S3Uri, response.AnimationFile),
		ImageFile:     fmt.Sprintf("%s%s", cfg.S3Uri, response.ImageFile),
		SoundFile:     fmt.Sprintf("%s%s", cfg.S3Uri, response.SoundFile),
		StartDate:     response.StartDate,
		EndDate:       response.EndDate,
		CreatedAt:     response.CreatedAt,
		UpdatedAt:     response.UpdatedAt,
		Gender:        response.Gender,
	}
	return routine, nil
}
