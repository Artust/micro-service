package routine_pos

import (
	pb "avatar/services/gateway/protos/pos"
	"context"

	"github.com/jinzhu/copier"
)

type UpdateRoutineInput struct {
	Id        int64
	Name      string `form:"name"`
	Detail    string `form:"detail"`
	StartDate string `form:"startDate"`
	EndDate   string `form:"endDate"`
	Gender    int64  `form:"gender" binding:"min=0,max=1"`
}

func UpdateRoutine(input *UpdateRoutineInput, posClient pb.POSClient) (*CreateRoutineOutput, error) {
	ctx := context.Background()
	oldRoutine, err := posClient.GetRoutine(ctx, &pb.GetByIdRequest{
		Id: input.Id,
	})
	if err != nil {
		return nil, err
	}
	updateRoutineRequest := &pb.UpdateRoutineRequest{}
	err = copier.Copy(updateRoutineRequest, oldRoutine)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(updateRoutineRequest, input)
	if err != nil {
		return nil, err
	}
	response, err := posClient.UpdateRoutine(context.Background(), updateRoutineRequest)
	if err != nil {
		return nil, err
	}
	result := &CreateRoutineOutput{
		Id:                       response.Id,
		Name:                     response.Name,
		Detail:                   response.Detail,
		AnimationFile:            response.AnimationFile,
		ImageFile:                response.ImageFile,
		SoundFile:                response.SoundFile,
		StartDate:                response.StartDate,
		EndDate:                  response.EndDate,
		PosId:                    response.PosId,
		ServiceTemplateId:        response.ServiceTemplateId,
		ServiceTemplateRoutineId: response.ServiceTemplateRoutineId,
		CategoryId:               response.CategoryId,
		CreatedAt:                response.CreatedAt,
		UpdatedAt:                response.UpdatedAt,
		Gender:                   response.Gender,
	}
	return result, nil
}
