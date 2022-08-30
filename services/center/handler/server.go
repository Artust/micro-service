package handler

import (
	"avatar/services/center/config"
	"avatar/services/center/domain/repository"
	pb "avatar/services/center/protos"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type CenterServer struct {
	pb.UnimplementedCenterServer
	cfg                               *config.Environment
	neo4jDriver                       neo4j.Driver
	serviceTemplateRepository         repository.ServiceTemplateRepository
	routineCategoryRepository         repository.RoutineCategoryRepository
	routineRepository                 repository.RoutineRepository
	avatarRepository                  repository.AvatarRepository
	serviceTemplateCategoryRepository repository.ServiceTemplateCategoryRepository
}

func CreateServer(
	cfg *config.Environment,
	neo4jDriver neo4j.Driver,
	serviceTemplateRepository repository.ServiceTemplateRepository,
	routineCategoryRepository repository.RoutineCategoryRepository,
	routineRepository repository.RoutineRepository,
	serviceTemplateCategoryRepository repository.ServiceTemplateCategoryRepository,
	avatarRepository repository.AvatarRepository,
) *CenterServer {
	return &CenterServer{
		UnimplementedCenterServer:         pb.UnimplementedCenterServer{},
		cfg:                               cfg,
		neo4jDriver:                       neo4jDriver,
		serviceTemplateRepository:         serviceTemplateRepository,
		routineCategoryRepository:         routineCategoryRepository,
		routineRepository:                 routineRepository,
		serviceTemplateCategoryRepository: serviceTemplateCategoryRepository,
		avatarRepository:                  avatarRepository,
	}
}
