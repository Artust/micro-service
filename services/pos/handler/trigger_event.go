package handler

import (
	pb "avatar/services/pos/protos"
	"avatar/services/pos/usecase/trigger_event"
	"context"
)

func (s *Server) TriggerEventOperatorSide(ctx context.Context, input *pb.TriggerEventOperatorSideRequest) (*pb.TriggerEventOperatorSideResponse, error) {
	return trigger_event.TriggerEventOperatorSide(ctx, s.neo4jDriver, s.routineRepository, input, s.config)
}
