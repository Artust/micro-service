package routine_pos

import (
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/pos"
	"context"
	"fmt"
)

type GetListRoutineByCategoryInput struct {
	PosId   int64  `form:"posId"`
	Page    int64  `form:"page"`
	PerPage int64  `form:"perPage"`
	EndDate string `form:"endDate"`
	Between int64  `form:"between"`
}

type GetListRoutineByCategoryOutput struct {
	Results []*Routine `json:"result"`
}

type Routine struct {
	Id      int64                  `json:"id"`
	Name    string                 `json:"name"`
	Routine []*CreateRoutineOutput `json:"routines"`
}

func GetListRoutineByCategory(
	input *GetListRoutineByCategoryInput,
	posClient pb.POSClient,
	cfg *config.Environment) (*GetListRoutineByCategoryOutput, error) {
	data := &pb.GetListRoutineByCategoryRequest{
		PosId:   input.PosId,
		Page:    input.Page,
		PerPage: input.PerPage,
		EndDate: input.EndDate,
		Between: input.Between,
	}
	response, err := posClient.GetListRoutineByCategory(context.Background(), data)
	if err != nil {
		return nil, err
	}
	var routine Routine
	var output GetListRoutineByCategoryOutput
	output.Results = make([]*Routine, 0)
	for _, val := range response.GetListRoutineSortedByCategoryResponse {
		routine.Routine = make([]*CreateRoutineOutput, 0)
		for _, response := range val.Routines {
			routine.Routine = append(routine.Routine, &CreateRoutineOutput{
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
			})
		}
		output.Results = append(output.Results, &Routine{
			Id:      val.Category.Id,
			Name:    val.Category.Name,
			Routine: routine.Routine,
		})
	}
	return &output, nil
}
