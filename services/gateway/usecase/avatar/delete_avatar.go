package avatar

import (
	pb "avatar/services/gateway/protos/center"
	"context"
)

type DeleteAvatarInput struct {
	Id int64
}

type DeleteAvatarOutPut struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func DeleteAvatar(input *DeleteAvatarInput, centerClient pb.CenterClient) (*DeleteAvatarOutPut, error) {
	data := &pb.DeleteByIdRequest{
		Id: input.Id,
	}
	response, err := centerClient.DeleteAvatar(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &DeleteAvatarOutPut{
		RowsAffected: response.RowsAffected,
	}
	return output, nil
}
