package routine_category_center

import (
	pb "avatar/services/gateway/protos/center"
	"context"
)

type GetRoutineCategoryInput struct {
	Id int64
}
type GetRoutineCategoryOutput struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func GetRoutineCategory(input *GetRoutineCategoryInput, centerClient pb.CenterClient) (*GetRoutineCategoryOutput, error) {
	ctx := context.Background()
	getRoutineCategoryInput := &pb.GetByIdRequest{
		Id: input.Id,
	}
	getRoutineCategoryResponse, err := centerClient.GetRoutineCategory(ctx, getRoutineCategoryInput)
	if err != nil {
		return nil, err
	}
	getRoutineCategoryOutput := &GetRoutineCategoryOutput{
		Id:        getRoutineCategoryResponse.Id,
		Name:      getRoutineCategoryResponse.Name,
		CreatedAt: getRoutineCategoryResponse.CreatedAt,
		UpdatedAt: getRoutineCategoryResponse.UpdatedAt,
	}

	return getRoutineCategoryOutput, nil
}
