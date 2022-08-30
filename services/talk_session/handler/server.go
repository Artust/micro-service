package handler

import (
	"avatar/services/talk_session/config"
	pb "avatar/services/talk_session/protos"
	dmRepository "avatar/services/talk_session/domain/repository"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Server struct {
	pb.UnimplementedTalkSessionServer
	config                     *config.Environment
	neo4jDriver                neo4j.Driver
	noteRepository             dmRepository.NoteRepository
}

func CreateServer(
	config *config.Environment,
	neo4jDriver neo4j.Driver,
	noteRepository dmRepository.NoteRepository,
) *Server {
	return &Server{
		config:                     config,
		noteRepository:             noteRepository,
		neo4jDriver:                neo4jDriver,
	}
}
