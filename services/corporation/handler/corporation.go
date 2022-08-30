package handler

import (
	pb "avatar/services/corporation/protos"
	"avatar/services/corporation/usecase/corporation"
	"context"
)

func (s *Server) CreateCorporation(ctx context.Context, input *pb.CreateCorporationRequest) (*pb.CreateCorporationResponse, error) {
	return corporation.Create(ctx, s.neo4jDriver, s.corporationRepository, input)
}

func (s *Server) GetCorporation(ctx context.Context, input *pb.GetCorporationRequest) (*pb.CreateCorporationResponse, error) {
	return corporation.GetById(ctx, s.neo4jDriver, s.corporationRepository, input)
}

func (s *Server) UpdateCorporation(ctx context.Context, input *pb.UpdateCorporationRequest) (*pb.CreateCorporationResponse, error) {
	return corporation.Update(ctx, s.neo4jDriver, s.corporationRepository, input)
}

func (s *Server) DeleteCorporation(ctx context.Context, input *pb.DeleteCorporationRequest) (*pb.DeleteCorporationResponse, error) {
	return corporation.Delete(ctx, s.neo4jDriver, s.corporationRepository, input)
}

func (s *Server) GetListCorporation(ctx context.Context, input *pb.GetListCorporationRequest) (*pb.GetListCorporationResponse, error) {
	return corporation.GetList(ctx, s.neo4jDriver, s.corporationRepository, input)
}
