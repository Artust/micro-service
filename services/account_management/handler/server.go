package handler

import (
	"avatar/services/account_management/config"
	"avatar/services/account_management/domain/mail"
	"avatar/services/account_management/domain/repository"
	pb "avatar/services/account_management/protos"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Server struct {
	accountRepository            repository.AccountRepository
	permissionRepository         repository.PermissionRepository
	accountRoleRepository        repository.AccountRoleRepository
	resetPasswordTokenRepository repository.ResetPasswordTokenRepository
	userActivityRepository       repository.UserActivityRepository
	mailClient                   mail.MailClient
	config                       *config.Environment
	neo4jDriver                  neo4j.Driver
	pb.UnimplementedAccountManagementServer
}

func CreateServer(
	accountRepository repository.AccountRepository,
	permissionRepository repository.PermissionRepository,
	accountRoleRepository repository.AccountRoleRepository,
	resetPasswordTokenRepository repository.ResetPasswordTokenRepository,
	userActivityRepository repository.UserActivityRepository,
	mailClient mail.MailClient,
	config *config.Environment,
	neo4jDriver neo4j.Driver,
) *Server {
	return &Server{
		accountRepository:            accountRepository,
		permissionRepository:         permissionRepository,
		accountRoleRepository:        accountRoleRepository,
		resetPasswordTokenRepository: resetPasswordTokenRepository,
		userActivityRepository:       userActivityRepository,
		mailClient:                   mailClient,
		config:                       config,
		neo4jDriver:                  neo4jDriver,
	}
}
