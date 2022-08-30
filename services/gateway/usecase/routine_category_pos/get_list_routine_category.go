package routine_category_pos

import (
	pb "avatar/services/gateway/protos/pos"
	"context"
)

type GetListRoutineCategoryInput struct {
	Page    int64 `form:"page"`
	PerPage int64 `form:"perPage"`
}

type GetListRoutineCategoryOutput struct {
	Results []*CreateRoutineCategoryOutput `json:"results"`
}

func GetListRoutineCategory(input *GetListRoutineCategoryInput, posClient pb.POSClient) (*GetListRoutineCategoryOutput, error) {
	data := &pb.GetListRoutineCategoryRequest{
		Page:    input.Page,
		PerPage: input.PerPage,
	}
	response, err := posClient.GetListRoutineCategory(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &GetListRoutineCategoryOutput{}
	output.Results = make([]*CreateRoutineCategoryOutput, 0)
	for _, v := range response.GetListRoutineCategoryResponse {
		routineCategory := &CreateRoutineCategoryOutput{
			Id:        v.Id,
			Name:      v.Name,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		output.Results = append(output.Results, routineCategory)
		routineCategory = nil
	}
	return output, nil
}
