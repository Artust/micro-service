package routine_center

import (
	"avatar/services/gateway/config"
	"avatar/services/gateway/pkg/util"
	pb "avatar/services/gateway/protos/center"
	"context"
	"fmt"
)

type GetListRoutineByCategoryInput struct {
	Page        int64  `form:"page"`
	PerPage     int64  `form:"perPage"`
	CategoryIds string `form:"categoryIds"`
	Ids         string `form:"ids"`
}

type GetListRoutineByCategoryOutput struct {
	Results []*Routine `json:"result"`
}

type Routine struct {
	Id          int64                        `json:"id"`
	Name        string                       `json:"name"`
	ListRoutine []*CreateRoutineCenterOutput `json:"routines"`
}

func GetListRoutineByCategory(
	input *GetListRoutineByCategoryInput,
	centerClient pb.CenterClient,
	cfg *config.Environment) (*GetListRoutineByCategoryOutput, error) {

	data := &pb.GetListRoutineByCategoryRequest{
		Page:        input.Page,
		PerPage:     input.PerPage,
		CategoryIds: util.GenerateIds(input.CategoryIds),
		Ids:         util.GenerateIds(input.Ids),
	}
	response, err := centerClient.GetListRoutineByCategory(context.Background(), data)
	if err != nil {
		return nil, err
	}
	var routine Routine
	var output GetListRoutineByCategoryOutput
	output.Results = make([]*Routine, 0)
	for _, val := range response.GetListRoutineByCategoryResponse {
		routine.ListRoutine = make([]*CreateRoutineCenterOutput, 0)
		for _, response := range val.Routine {
			routine.ListRoutine = append(routine.ListRoutine, &CreateRoutineCenterOutput{
				Id:            response.Id,
				Name:          response.Name,
				Detail:        response.Detail,
				AnimationFile: fmt.Sprintf("%s%s", cfg.S3Uri, response.AnimationFile),
				ImageFile:     fmt.Sprintf("%s%s", cfg.S3Uri, response.ImageFile),
				SoundFile:     fmt.Sprintf("%s%s", cfg.S3Uri, response.SoundFile),
				StartDate:     response.StartDate,
				EndDate:       response.EndDate,
				CategoryId:    response.CategoryId,
				Gender:        response.Gender,
				CreatedAt:     response.CreatedAt,
				UpdatedAt:     response.UpdatedAt,
			})
		}
		output.Results = append(output.Results, &Routine{
			Id:          val.Category.Id,
			Name:        val.Category.Name,
			ListRoutine: routine.ListRoutine,
		})
	}
	return &output, nil
}
