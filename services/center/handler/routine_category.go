package handler

import (
	pb "avatar/services/center/protos"
	"avatar/services/center/usecase/routine_category"
	"context"
)

func (c *CenterServer) CreateRoutineCategory(
	ctx context.Context,
	input *pb.CreateRoutineCategoryRequest,
) (*pb.CreateRoutineCategoryResponse, error) {
	return routine_category.Create(ctx, c.neo4jDriver, c.routineCategoryRepository, input)
}

func (c *CenterServer) GetRoutineCategory(
	ctx context.Context,
	input *pb.GetByIdRequest,
) (*pb.CreateRoutineCategoryResponse, error) {
	return routine_category.GetById(ctx, c.neo4jDriver, c.routineCategoryRepository, input)
}

func (c *CenterServer) GetListRoutineCategory(
	ctx context.Context,
	input *pb.GetListRoutineCategoryRequest,
) (*pb.GetListRoutineCategoryResponse, error) {
	return routine_category.GetList(ctx, c.neo4jDriver, c.routineCategoryRepository, input)
}

func (c *CenterServer) UpdateRoutineCategory(
	ctx context.Context,
	input *pb.UpdateRoutineCategoryRequest,
) (*pb.CreateRoutineCategoryResponse, error) {
	return routine_category.Update(ctx, c.neo4jDriver, c.routineCategoryRepository, input)
}

func (c *CenterServer) DeleteRoutineCategory(
	ctx context.Context,
	input *pb.DeleteByIdRequest,
) (*pb.DeleteResponse, error) {
	return routine_category.Delete(ctx, c.neo4jDriver, c.routineCategoryRepository, input)
}
