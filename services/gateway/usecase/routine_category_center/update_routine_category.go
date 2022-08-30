package routine_category_center

import (
	pb "avatar/services/gateway/protos/center"
	"context"
)

type UpdateRoutineCategoryInput struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type UpdateRoutineCategoryOutput struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func UpdateRoutineCategory(input *UpdateRoutineCategoryInput, centerClient pb.CenterClient) (*UpdateRoutineCategoryOutput, error) {
	ctx := context.Background()

	updateRoutineCategoryInput := &pb.UpdateRoutineCategoryRequest{
		Id:   input.Id,
		Name: input.Name,
	}
	updateRoutineCategoryResponse, err := centerClient.UpdateRoutineCategory(ctx, updateRoutineCategoryInput)
	if err != nil {
		return nil, err
	}
	updateRoutineCategoryOutput := &UpdateRoutineCategoryOutput{
		Id:        updateRoutineCategoryResponse.Id,
		Name:      updateRoutineCategoryResponse.Name,
		CreatedAt: updateRoutineCategoryResponse.CreatedAt,
		UpdatedAt: updateRoutineCategoryResponse.UpdatedAt,
	}

	return updateRoutineCategoryOutput, nil
}
