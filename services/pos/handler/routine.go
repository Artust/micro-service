package handler

import (
	pb "avatar/services/pos/protos"
	"avatar/services/pos/usecase/routine"
	"context"
)

func (s *Server) CreateRoutine(ctx context.Context, input *pb.CreateRoutineRequest) (*pb.CreateRoutineResponse, error) {
	return routine.Create(ctx, s.neo4jDriver, s.routineRepository, input)
}

func (s *Server) CreateManyRoutine(ctx context.Context, input *pb.CreateManyRoutineRequest) (*pb.CreateManyRoutineResponse, error) {
	return routine.CreateMany(ctx, s.neo4jDriver, s.routineRepository, input)
}

func (s *Server) GetRoutine(ctx context.Context, input *pb.GetByIdRequest) (*pb.CreateRoutineResponse, error) {
	return routine.GetById(ctx, s.neo4jDriver, s.routineRepository, input)
}

func (s *Server) GetListRoutine(ctx context.Context, input *pb.GetListRoutineRequest) (*pb.GetListRoutineResponse, error) {
	return routine.GetList(ctx, s.neo4jDriver, s.routineRepository, input)
}

func (s *Server) UpdateRoutine(ctx context.Context, input *pb.UpdateRoutineRequest) (*pb.CreateRoutineResponse, error) {
	return routine.Update(ctx, s.neo4jDriver, s.routineRepository, input)
}

func (s *Server) DeleteRoutine(ctx context.Context, input *pb.DeleteRoutineRequest) (*pb.DeleteResponse, error) {
	return routine.Delete(ctx, s.neo4jDriver, s.routineRepository, input)
}

func (s *Server) GetListRoutineByCategory(ctx context.Context, input *pb.GetListRoutineByCategoryRequest) (*pb.GetListRoutineByCategoryResponse, error) {
	return routine.GetListByCategory(ctx, s.neo4jDriver, s.routineRepository, s.routineCategoryRepository, input)
}
