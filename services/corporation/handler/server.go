package handler

import (
	"avatar/services/corporation/config"
	"avatar/services/corporation/domain/repository"
	pb "avatar/services/corporation/protos"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Server struct {
	pb.UnimplementedCorporationServer
	config                *config.Environment
	neo4jDriver           neo4j.Driver
	deviceRepository      repository.DeviceRepository
	shopRepository        repository.ShopRepository
	corporationRepository repository.CorporationRepository
	centerRepository      repository.CenterRepository
}

func CreateServer(
	config *config.Environment,
	neo4jDriver neo4j.Driver,
	deviceRepository repository.DeviceRepository,
	shopRepository repository.ShopRepository,
	corporationRepository repository.CorporationRepository,
	centerRepository repository.CenterRepository,
) *Server {
	return &Server{
		config:                config,
		neo4jDriver:           neo4jDriver,
		deviceRepository:      deviceRepository,
		shopRepository:        shopRepository,
		corporationRepository: corporationRepository,
		centerRepository:      centerRepository,
	}
}
