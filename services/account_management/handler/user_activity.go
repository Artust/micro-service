package handler

import (
	pb "avatar/services/account_management/protos"
	"avatar/services/account_management/usecase/user_activity"
	"context"
)

func (s *Server) GetUserActivity(ctx context.Context, input *pb.AccountId) (*pb.UserActivityList, error) {
	return user_activity.Get(ctx, s.neo4jDriver, s.userActivityRepository, input)
}

func (s *Server) SaveUserActivity(ctx context.Context, input *pb.UserActivity) (*pb.UserActivity, error) {
	return user_activity.Save(ctx, s.neo4jDriver, s.userActivityRepository, input)
}
