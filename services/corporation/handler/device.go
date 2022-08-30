package handler

import (
	pb "avatar/services/corporation/protos"
	"avatar/services/corporation/usecase/device"
	"context"
)

func (s *Server) CreateDevice(ctx context.Context, input *pb.CreateDeviceRequest) (*pb.CreateDeviceResponse, error) {
	return device.Create(ctx, s.neo4jDriver, s.deviceRepository, input)
}

func (s *Server) GetDevice(ctx context.Context, input *pb.GetDeviceRequest) (*pb.CreateDeviceResponse, error) {
	return device.GetById(ctx, s.neo4jDriver, s.deviceRepository, input)
}

func (s *Server) UpdateDevice(ctx context.Context, input *pb.UpdateDeviceRequest) (*pb.CreateDeviceResponse, error) {
	return device.Update(ctx, s.neo4jDriver, s.deviceRepository, input)
}

func (s *Server) DeleteDevice(ctx context.Context, input *pb.DeleteDeviceRequest) (*pb.DeleteDeviceResponse, error) {
	return device.Delete(ctx, s.neo4jDriver, s.deviceRepository, input)
}

func (s *Server) GetListDevice(ctx context.Context, input *pb.GetListDeviceRequest) (*pb.GetListDeviceResponse, error) {
	return device.GetList(ctx, s.neo4jDriver, s.deviceRepository, input)
}
