package routine_center

import (
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/center"
	"context"
	"fmt"
	// log "github.com/sirupsen/logrus"
)

type UpdateRoutineInput struct {
	Id            int64
	Name          string `json:"name"`
	Detail        string `json:"detail"`
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
	AnimationFile string `json:"animationFile"`
	ImageFile     string `json:"imageFile"`
	SoundFile     string `json:"soundFile"`
	CategoryId    int64  `json:"categoryId"`
	Gender        int64  `json:"gender" binding:"min=0,max=1"`
}

func UpdateRoutine(
	input *UpdateRoutineInput,
	centerClient pb.CenterClient,
	cfg *config.Environment,
) (*CreateRoutineCenterOutput, error) {
	response, err := centerClient.UpdateRoutine(context.Background(), &pb.UpdateRoutineRequest{
		Id:            input.Id,
		Name:          input.Name,
		Detail:        input.Detail,
		AnimationFile: input.AnimationFile,
		ImageFile:     input.ImageFile,
		SoundFile:     input.SoundFile,
		StartDate:     input.StartDate,
		EndDate:       input.EndDate,
		CategoryId:    input.CategoryId,
		Gender:        input.Gender,
	})
	if err != nil {
		return nil, err
	}
	result := &CreateRoutineCenterOutput{
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
		CategoryId:    response.CategoryId,
		Gender:        response.Gender,
	}
	return result, nil
}
