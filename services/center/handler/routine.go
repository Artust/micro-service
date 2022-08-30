package handler

import (
	pb "avatar/services/center/protos"
	"avatar/services/center/usecase/routine"
	"context"
)

func (c *CenterServer) GetListRoutine(ctx context.Context, input *pb.GetListRoutineRequest) (*pb.GetListRoutineResponse, error) {
	return routine.GetList(ctx, c.neo4jDriver, c.routineRepository, input)
}

func (c *CenterServer) GetRoutine(ctx context.Context, input *pb.GetByIdRequest) (*pb.CreateRoutineResponse, error) {
	return routine.GetById(ctx, c.neo4jDriver, c.routineRepository, input)
}

func (c *CenterServer) CreateRoutine(ctx context.Context, input *pb.CreateRoutineRequest) (*pb.CreateRoutineResponse, error) {
	return routine.Create(ctx, c.neo4jDriver, c.routineRepository, input)
}

func (c *CenterServer) DeleteRoutine(ctx context.Context, input *pb.DeleteByIdRequest) (*pb.DeleteResponse, error) {
	return routine.Delete(ctx, c.neo4jDriver, c.routineRepository, input)
}

func (c *CenterServer) UpdateRoutine(ctx context.Context, input *pb.UpdateRoutineRequest) (*pb.CreateRoutineResponse, error) {
	return routine.Update(ctx, c.neo4jDriver, c.routineRepository, input)
}

func (s *CenterServer) GetListRoutineByCategory(ctx context.Context, input *pb.GetListRoutineByCategoryRequest) (*pb.GetListRoutineByCategoryResponse, error) {
	return routine.GetListByCategory(ctx, s.neo4jDriver, s.routineRepository, s.routineCategoryRepository, input)
}
