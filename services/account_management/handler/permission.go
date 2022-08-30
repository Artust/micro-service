package handler

import (
	pb "avatar/services/account_management/protos"
	"avatar/services/account_management/usecase/permission"
	"context"
)

func (s *Server) GetPermission(ctx context.Context, input *pb.Id) (*pb.Permission, error) {
	return permission.GetById(ctx, s.neo4jDriver, s.permissionRepository, input)
}

func (s *Server) GetListPermission(ctx context.Context, input *pb.Empty) (*pb.PermissionList, error) {
	return permission.GetList(ctx, s.neo4jDriver, s.permissionRepository)
}

func (s *Server) CreatePermission(ctx context.Context, input *pb.CreatePermissionRequest) (*pb.Permission, error) {
	return permission.Create(ctx, s.neo4jDriver, s.permissionRepository, input)
}

func (s *Server) UpdatePermission(ctx context.Context, input *pb.UpdatePermissionRequest) (*pb.Permission, error) {
	return permission.Update(ctx, s.neo4jDriver, s.permissionRepository, input)
}

func (s *Server) DeletePermission(ctx context.Context, input *pb.DeletePermissionRequest) (*pb.Empty, error) {
	return permission.Delete(ctx, s.neo4jDriver, s.permissionRepository, input)
}
