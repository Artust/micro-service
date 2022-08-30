package handler

import (
	pb "avatar/services/pos/protos"
	"avatar/services/pos/usecase/routine_category"
	"context"
)

func (s *Server) CreateRoutineCategory(ctx context.Context, input *pb.CreateRoutineCategoryRequest) (*pb.CreateRoutineCategoryResponse, error) {
	return routine_category.Create(ctx, s.neo4jDriver, s.routineCategoryRepository, input)
}

func (s *Server) GetListRoutineCategory(ctx context.Context, input *pb.GetListRoutineCategoryRequest) (*pb.GetListRoutineCategoryResponse, error) {
	return routine_category.GetList(ctx, s.neo4jDriver, s.routineCategoryRepository, input)
}

func (s *Server) GetRoutineCategory(ctx context.Context, input *pb.GetByIdRequest) (*pb.CreateRoutineCategoryResponse, error) {
	return routine_category.GetById(ctx, s.neo4jDriver, s.routineCategoryRepository, input)
}

func (s *Server) UpdateRoutineCategory(ctx context.Context, input *pb.UpdateRoutineCategoryRequest) (*pb.CreateRoutineCategoryResponse, error) {
	return routine_category.Update(ctx, s.neo4jDriver, s.routineCategoryRepository, input)
}

func (s *Server) DeleteRoutineCategory(ctx context.Context, input *pb.DeleteByIdRequest) (*pb.DeleteResponse, error) {
	return routine_category.Delete(ctx, s.neo4jDriver, s.routineCategoryRepository, input)
}
