package handler

import (
	pb "avatar/services/pos/protos"
	"avatar/services/pos/usecase/ip_camera"
	"context"
)

func (s *Server) CreateIpCamera(ctx context.Context, input *pb.CreateIpCameraRequest) (*pb.CreateIpCameraResponse, error) {
	return ip_camera.Create(ctx, s.neo4jDriver, s.ipCameraRepository, input)
}

func (s *Server) GetIpCamera(ctx context.Context, input *pb.GetByIdRequest) (*pb.CreateIpCameraResponse, error) {
	return ip_camera.GetById(ctx, s.neo4jDriver, s.ipCameraRepository, input)
}

func (s *Server) UpdateIpCamera(ctx context.Context, input *pb.UpdateIpCameraRequest) (*pb.CreateIpCameraResponse, error) {
	return ip_camera.Update(ctx, s.neo4jDriver, s.ipCameraRepository, input)
}

func (s *Server) DeleteIpCamera(ctx context.Context, input *pb.DeleteByIdRequest) (*pb.DeleteResponse, error) {
	return ip_camera.Delete(ctx, s.neo4jDriver, s.ipCameraRepository, input)
}

func (s *Server) GetListIpCamera(ctx context.Context, input *pb.GetListIpCameraRequest) (*pb.GetListIpCameraResponse, error) {
	return ip_camera.GetList(ctx, s.neo4jDriver, s.ipCameraRepository, input)
}
