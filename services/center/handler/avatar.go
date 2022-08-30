package handler

import (
	pb "avatar/services/center/protos"
	"avatar/services/center/usecase/avatar"
	"context"
)

func (c *CenterServer) CreateAvatar(ctx context.Context, input *pb.CreateAvatarRequest) (*pb.CreateAvatarResponse, error) {
	return avatar.Create(ctx, c.neo4jDriver, c.avatarRepository, input)
}

func (c *CenterServer) GetAvatar(ctx context.Context, input *pb.GetByIdRequest) (*pb.CreateAvatarResponse, error) {
	return avatar.GetById(ctx, c.neo4jDriver, c.avatarRepository, input)
}

func (c *CenterServer) GetListAvatar(ctx context.Context, input *pb.GetListAvatarRequest) (*pb.GetListAvatarResponse, error) {
	return avatar.GetList(ctx, c.neo4jDriver, c.avatarRepository, input)
}

func (c *CenterServer) UpdateAvatar(ctx context.Context, input *pb.CreateAvatarRequest) (*pb.CreateAvatarResponse, error) {
	return avatar.Update(ctx, c.neo4jDriver, c.avatarRepository, input)
}

func (c *CenterServer) DeleteAvatar(ctx context.Context, input *pb.DeleteByIdRequest) (*pb.DeleteResponse, error) {
	return avatar.Delete(ctx, c.neo4jDriver, c.avatarRepository, input)
}
