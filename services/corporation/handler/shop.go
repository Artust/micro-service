package handler

import (
	pb "avatar/services/corporation/protos"
	"avatar/services/corporation/usecase/shop"
	"context"
)

func (s *Server) CreateShop(ctx context.Context, input *pb.CreateShopRequest) (*pb.CreateShopResponse, error) {
	return shop.Create(ctx, s.neo4jDriver, s.shopRepository, input)
}

func (s *Server) GetShop(ctx context.Context, input *pb.GetShopRequest) (*pb.CreateShopResponse, error) {
	return shop.GetById(ctx, s.neo4jDriver, s.shopRepository, input)
}

func (s *Server) UpdateShop(ctx context.Context, input *pb.UpdateShopRequest) (*pb.CreateShopResponse, error) {
	return shop.Update(ctx, s.neo4jDriver, s.shopRepository, input)
}

func (s *Server) DeleteShop(ctx context.Context, input *pb.DeleteShopRequest) (*pb.DeleteShopResponse, error) {
	return shop.Delete(ctx, s.neo4jDriver, s.shopRepository, input)
}

func (s *Server) GetListShop(ctx context.Context, input *pb.GetListShopRequest) (*pb.GetListShopResponse, error) {
	return shop.GetList(ctx, s.neo4jDriver, s.shopRepository, input)
}

