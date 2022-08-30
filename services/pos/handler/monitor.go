package handler

import (
	pb "avatar/services/pos/protos"
	"avatar/services/pos/usecase/monitor"
	"context"
)

func (s *Server) CreateMonitor(ctx context.Context, input *pb.CreateMonitorRequest) (*pb.CreateMonitorResponse, error) {
	return monitor.Create(ctx, s.neo4jDriver, s.monitorRepository, input)
}

func (s *Server) GetMonitor(ctx context.Context, input *pb.GetByIdRequest) (*pb.CreateMonitorResponse, error) {
	return monitor.GetById(ctx, s.neo4jDriver, s.monitorRepository, input)
}

func (s *Server) UpdateMonitor(ctx context.Context, input *pb.UpdateMonitorRequest) (*pb.CreateMonitorResponse, error) {
	return monitor.Update(ctx, s.neo4jDriver, s.monitorRepository, input)
}

func (s *Server) DeleteMonitor(ctx context.Context, input *pb.DeleteByIdRequest) (*pb.DeleteResponse, error) {
	return monitor.Delete(ctx, s.neo4jDriver, s.monitorRepository, input)
}

func (s *Server) GetListMonitor(ctx context.Context, input *pb.GetListMonitorRequest) (*pb.GetListMonitorResponse, error) {
	return monitor.GetList(ctx, s.neo4jDriver, s.monitorRepository, input)
}
