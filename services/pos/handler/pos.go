package handler

import (
	pb "avatar/services/pos/protos"
	"avatar/services/pos/usecase/pos"
	"context"
)

func (s *Server) CreatePos(ctx context.Context, input *pb.CreatePosRequest) (*pb.CreatePosResponse, error) {
	return pos.Create(ctx, s.neo4jDriver, s.posRepository, input)
}

func (s *Server) GetPos(ctx context.Context, input *pb.GetByIdRequest) (*pb.CreatePosResponse, error) {
	return pos.GetById(ctx, s.neo4jDriver, s.posRepository, input)
}

func (s *Server) UpdatePos(ctx context.Context, input *pb.UpdatePosRequest) (*pb.CreatePosResponse, error) {
	return pos.Update(ctx, s.neo4jDriver, s.posRepository, input)
}

func (s *Server) DeletePos(ctx context.Context, input *pb.DeleteByIdRequest) (*pb.DeleteResponse, error) {
	return pos.Delete(ctx, s.neo4jDriver, s.posRepository, input)
}

func (s *Server) GetListPos(ctx context.Context, input *pb.GetListPosRequest) (*pb.GetListPosResponse, error) {
	return pos.GetList(ctx, s.neo4jDriver, s.posRepository, input)
}
