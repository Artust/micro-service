package routine_category_center

import (
	pb "avatar/services/gateway/protos/center"
	"context"
)

type CreateRoutineCategoryInput struct {
	Name string `json:"name" binding:"required"`
}
type CreateRoutineCategoryOutput struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func CreateRoutineCategory(input *CreateRoutineCategoryInput, centerClient pb.CenterClient) (*CreateRoutineCategoryOutput, error) {
	ctx := context.Background()
	createRoutineCategoryInput := &pb.CreateRoutineCategoryRequest{
		Name: input.Name,
	}
	createRoutineCategoryResponse, err := centerClient.CreateRoutineCategory(ctx, createRoutineCategoryInput)
	if err != nil {
		return nil, err
	}
	createRoutineCategoryOutput := &CreateRoutineCategoryOutput{
		Id:        createRoutineCategoryResponse.Id,
		Name:      createRoutineCategoryResponse.Name,
		CreatedAt: createRoutineCategoryResponse.CreatedAt,
		UpdatedAt: createRoutineCategoryResponse.UpdatedAt,
	}

	return createRoutineCategoryOutput, nil
}
