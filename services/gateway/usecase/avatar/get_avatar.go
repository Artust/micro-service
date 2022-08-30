package avatar

import (
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/center"
	"context"
	"fmt"
)

type GetAvatarInput struct {
	Id int64
}

func GetAvatar(
	input *GetAvatarInput,
	centerClient pb.CenterClient,
	cfg *config.Environment,
) (*CreateAvatarOutPut, error) {
	data := &pb.GetByIdRequest{
		Id: input.Id,
	}
	response, err := centerClient.GetAvatar(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &CreateAvatarOutPut{
		Id:        response.Id,
		Name:      response.Name,
		Detail:    response.Detail,
		Image:     fmt.Sprintf("%s%s", cfg.S3Uri, response.Image),
		Vrm:       fmt.Sprintf("%s%s", cfg.S3Uri, response.Vrm),
		StartDate: response.StartDate,
		EndDate:   response.EndDate,
		Gender:    response.Gender,
		Version:   response.Version,
		Exporter:  response.Exporter,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	}
	return output, nil
}
