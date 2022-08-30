package routine_category_center

import (
	pb "avatar/services/gateway/protos/center"
	"context"
)

type GetListRoutineCategoryInput struct {
	Page    int64 `form:"page"`
	PerPage int64 `form:"perPage"`
}
type GetListRoutineCategoryOutput struct {
	ListRoutineCategory []*RoutineCategory `json:"results"`
}
type RoutineCategory struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func GetListRoutineCategory(input *GetListRoutineCategoryInput, centerClient pb.CenterClient) (*GetListRoutineCategoryOutput, error) {
	ctx := context.Background()

	getListRoutineCategoryInput := &pb.GetListRoutineCategoryRequest{
		Page:    int64(input.Page),
		PerPage: int64(input.PerPage),
	}
	getListRoutineCategoryResponse, err := centerClient.GetListRoutineCategory(ctx, getListRoutineCategoryInput)
	if err != nil {
		return nil, err
	}
	getListRoutineCategoryOutput := &GetListRoutineCategoryOutput{}
	getListRoutineCategoryOutput.ListRoutineCategory = make([]*RoutineCategory, 0)
	for _, v := range getListRoutineCategoryResponse.GetListRoutineCategoryResponse {
		routineCategory := &RoutineCategory{
			Id:        v.Id,
			Name:      v.Name,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		getListRoutineCategoryOutput.ListRoutineCategory = append(getListRoutineCategoryOutput.ListRoutineCategory, routineCategory)
		routineCategory = nil
	}
	return getListRoutineCategoryOutput, nil
}
