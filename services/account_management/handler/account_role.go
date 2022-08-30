package handler

import (
	pb "avatar/services/account_management/protos"
	"avatar/services/account_management/usecase/account_role"
	"context"
)

func (s *Server) GetAccountRole(ctx context.Context, input *pb.Id) (*pb.AccountRole, error) {
	return account_role.GetById(ctx, s.neo4jDriver, s.accountRoleRepository, input)
}

func (s *Server) GetListAccountRole(ctx context.Context, input *pb.Empty) (*pb.AccountRoleList, error) {
	return account_role.GetList(ctx, s.neo4jDriver, s.accountRoleRepository)
}

func (s *Server) CreateAccountRole(ctx context.Context, input *pb.CreateAccountRoleRequest) (*pb.AccountRole, error) {
	return account_role.Create(ctx, s.neo4jDriver, s.accountRoleRepository, input)
}

func (s *Server) UpdateAccountRole(ctx context.Context, input *pb.UpdateAccountRoleRequest) (*pb.AccountRole, error) {
	return account_role.Update(ctx, s.neo4jDriver, s.accountRoleRepository, input)
}

func (s *Server) DeleteAccountRole(ctx context.Context, input *pb.DeleteAccountRoleRequest) (*pb.Empty, error) {
	return account_role.Delete(ctx, s.neo4jDriver, s.accountRoleRepository, input)
}
