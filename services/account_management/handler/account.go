package handler

import (
	"avatar/services/account_management/pkg/validate"
	pb "avatar/services/account_management/protos"
	"avatar/services/account_management/usecase/account"
	"context"
	"errors"
)

func (s *Server) Login(ctx context.Context, input *pb.LoginRequest) (*pb.LoginResponse, error) {
	if !validate.ValidEmail(input.Email) && !validate.ValidPassword(input.Password) {
		return nil, errors.New("invalid input")
	}
	return account.Login(ctx, s.config, s.neo4jDriver, s.accountRepository, input)
}

func (s *Server) ForgotPassword(ctx context.Context, input *pb.ForgotPasswordRequest) (*pb.Empty, error) {
	if !validate.ValidEmail(input.Email) {
		return nil, errors.New("invalid input")
	}
	return account.ForgotPassword(
		ctx,
		s.config,
		s.neo4jDriver,
		s.accountRepository,
		s.resetPasswordTokenRepository,
		s.mailClient,
		input,
	)
}

func (s *Server) ResetPassword(ctx context.Context, input *pb.ResetPasswordRequest) (*pb.Empty, error) {
	return account.ResetPassword(ctx, s.neo4jDriver, s.accountRepository, s.resetPasswordTokenRepository, input)
}

func (s *Server) GetAccount(ctx context.Context, input *pb.Id) (*pb.Account, error) {
	return account.GetById(ctx, s.neo4jDriver, s.accountRepository, input)
}

func (s *Server) GetListAccount(ctx context.Context, input *pb.GetListAccountRequest) (*pb.AccountList, error) {
	return account.GetList(ctx, s.neo4jDriver, s.accountRepository, input)
}

func (s *Server) CreateAccount(ctx context.Context, input *pb.CreateAccountRequest) (*pb.Account, error) {
	if !validate.ValidEmail(input.Email) {
		return nil, errors.New("invalid input")
	}
	return account.Create(
		ctx,
		s.config,
		s.neo4jDriver,
		s.accountRepository,
		s.resetPasswordTokenRepository,
		s.mailClient,
		input,
	)
}

func (s *Server) UpdateAccount(ctx context.Context, input *pb.UpdateAccountRequest) (*pb.Account, error) {
	return account.Update(ctx, s.neo4jDriver, s.accountRepository, input)
}

func (s *Server) ChangePassword(ctx context.Context, input *pb.ChangePasswordRequest) (*pb.Empty, error) {
	return account.ChangePassword(ctx, s.neo4jDriver, s.accountRepository, input)
}

func (s *Server) DeactiveAccount(ctx context.Context, input *pb.DeactiveAccountRequest) (*pb.Empty, error) {
	return account.Deactive(ctx, s.neo4jDriver, s.accountRepository, input)
}

func (s *Server) ActiveAccount(ctx context.Context, input *pb.ActiveAccountRequest) (*pb.Empty, error) {
	return account.Active(ctx, s.neo4jDriver, s.accountRepository, input)
}
