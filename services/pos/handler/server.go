package handler

import (
	"avatar/services/pos/config"
	"avatar/services/pos/domain/repository"
	pb "avatar/services/pos/protos"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Server struct {
	pb.UnimplementedPOSServer
	config                    *config.Environment
	neo4jDriver               neo4j.Driver
	monitorRepository         repository.MonitorRepository
	ipCameraRepository        repository.IpCameraRepository
	posRepository             repository.PosRepository
	routineRepository         repository.RoutineRepository
	routineCategoryRepository repository.RoutineCategoryRepository
}

func CreateServer(
	config *config.Environment,
	neo4jDriver neo4j.Driver,
	posRepository repository.PosRepository,
	monitorRepository repository.MonitorRepository,
	ipCameraRepository repository.IpCameraRepository,
	routineRepository repository.RoutineRepository,
	routineCategoryRepository repository.RoutineCategoryRepository,
) *Server {
	return &Server{
		config:                    config,
		posRepository:             posRepository,
		monitorRepository:         monitorRepository,
		ipCameraRepository:        ipCameraRepository,
		routineRepository:         routineRepository,
		routineCategoryRepository: routineCategoryRepository,
		neo4jDriver:               neo4jDriver,
	}
}
