package routine_center

import (
	"avatar/services/gateway/config"
	"avatar/services/gateway/pkg/util"
	pb "avatar/services/gateway/protos/center"
	"context"
	"fmt"
)

type GetListRoutineInput struct {
	CategoryId int64  `form:"categoryId"`
	Page       int64  `form:"page"`
	PerPage    int64  `form:"perPage"`
	Gender     int64  `form:"gender" binding:"min=0,max=1"`
	Ids        string `form:"ids"`
}
type GetListRoutineOutPut struct {
	ListRoutine []*CreateRoutineCenterOutput `json:"results"`
}

func GetListRoutineCenter(
	input *GetListRoutineInput,
	centerClient pb.CenterClient,
	cfg *config.Environment,
) (*GetListRoutineOutPut, error) {
	listRoutineInput := &pb.GetListRoutineRequest{
		Page:       input.Page,
		PerPage:    input.PerPage,
		CategoryId: input.CategoryId,
		Gender:     input.Gender,
		Ids:        util.GenerateIds(input.Ids),
	}
	getListRoutineResponse, err := centerClient.GetListRoutine(context.Background(), listRoutineInput)
	if err != nil {
		return nil, err
	}

	getRoutineOutPut := &GetListRoutineOutPut{}
	getRoutineOutPut.ListRoutine = make([]*CreateRoutineCenterOutput, 0)
	for _, data := range getListRoutineResponse.ListRoutine {
		routine := &CreateRoutineCenterOutput{
			Id:            data.Id,
			Name:          data.Name,
			Detail:        data.Detail,
			AnimationFile: fmt.Sprintf("%s%s", cfg.S3Uri, data.AnimationFile),
			ImageFile:     fmt.Sprintf("%s%s", cfg.S3Uri, data.ImageFile),
			SoundFile:     fmt.Sprintf("%s%s", cfg.S3Uri, data.SoundFile),
			CategoryId:    data.CategoryId,
			StartDate:     data.StartDate,
			EndDate:       data.EndDate,
			Gender:        data.Gender,
			CreatedAt:     data.CreatedAt,
			UpdatedAt:     data.UpdatedAt,
		}
		getRoutineOutPut.ListRoutine = append(getRoutineOutPut.ListRoutine, routine)
		routine = nil
	}

	return getRoutineOutPut, nil
}
