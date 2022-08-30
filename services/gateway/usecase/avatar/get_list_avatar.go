package avatar

import (
	"avatar/services/gateway/config"
	"avatar/services/gateway/pkg/util"
	pb "avatar/services/gateway/protos/center"
	"context"
	"fmt"
	"log"
)

type GetListAvatarInput struct {
	Gender  int64  `form:"gender"`
	Page    int64  `form:"page"`
	PerPage int64  `form:"perPage"`
	Ids     string `form:"ids"`
}

type GetListAvatarOutPut struct {
	Results []*CreateAvatarOutPut `json:"results"`
}

func GetListAvatar(
	input *GetListAvatarInput,
	centerClient pb.CenterClient,
	cfg *config.Environment,
) (*GetListAvatarOutPut, error) {
	log.Println(util.GenerateIds(input.Ids))
	data := &pb.GetListAvatarRequest{
		Gender:  input.Gender,
		Page:    input.Page,
		PerPage: input.PerPage,
		Ids:     util.GenerateIds(input.Ids),
	}
	response, err := centerClient.GetListAvatar(context.Background(), data)
	if err != nil {
		return nil, err
	}
	var output GetListAvatarOutPut
	output.Results = make([]*CreateAvatarOutPut, 0)
	for _, response := range response.Avatars {
		output.Results = append(output.Results, &CreateAvatarOutPut{
			Id:        response.Id,
			Name:      response.Name,
			Detail:    response.Detail,
			Image:     fmt.Sprintf("%s%s", cfg.S3Uri, response.Image),
			Vrm:       fmt.Sprintf("%s%s", cfg.S3Uri, response.Vrm),
			StartDate: response.StartDate,
			EndDate:   response.EndDate,
			Version:   response.Version,
			Exporter:  response.Exporter,
			Gender:    response.Gender,
			CreatedAt: response.CreatedAt,
			UpdatedAt: response.UpdatedAt,
		})
	}
	return &output, nil
}
