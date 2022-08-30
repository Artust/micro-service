package handler

import (
	pb "avatar/services/corporation/protos"
	"avatar/services/corporation/usecase/center"
	"context"
)

func (s *Server) CreateCenter(ctx context.Context, input *pb.CreateCenterRequest) (*pb.CreateCenterResponse, error) {
	return center.Create(ctx, s.neo4jDriver, s.centerRepository, input)
}

func (s *Server) GetCenter(ctx context.Context, input *pb.GetCenterRequest) (*pb.CreateCenterResponse, error) {
	return center.GetById(ctx, s.neo4jDriver, s.centerRepository, input)
}

func (s *Server) UpdateCenter(ctx context.Context, input *pb.UpdateCenterRequest) (*pb.CreateCenterResponse, error) {
	return center.Update(ctx, s.neo4jDriver, s.centerRepository, input)
}

func (s *Server) DeleteCenter(ctx context.Context, input *pb.DeleteCenterRequest) (*pb.DeleteCenterResponse, error) {
	return center.Delete(ctx, s.neo4jDriver, s.centerRepository, input)
}

func (s *Server) GetListCenter(ctx context.Context, input *pb.GetListCenterRequest) (*pb.GetListCenterResponse, error) {
	return center.GetList(ctx, s.neo4jDriver, s.centerRepository, input)
}
