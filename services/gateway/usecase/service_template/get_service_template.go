package service_template

import (
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/center"
	"context"
	"fmt"
)

type GetServiceTemplateInput struct {
	ID int64 `json:"id"`
}

func GetServiceTemplate(
	input *GetServiceTemplateInput,
	centerClient pb.CenterClient,
	cfg *config.Environment,
) (*CreateServiceTemplateOutput, error) {
	ctx := context.Background()
	in := &pb.GetByIdRequest{
		Id: input.ID,
	}
	response, err := centerClient.GetServiceTemplate(ctx, in)
	if err != nil {
		return nil, err
	}
	backgroundName := []string{}
	for _, background := range response.Backgrounds {
		backgroundName = append(backgroundName, fmt.Sprintf("%s/%s/%s", cfg.S3Uri, config.BackgroundBucketName, background))
	}
	output := &CreateServiceTemplateOutput{
		Id:                      response.Id,
		Name:                    response.Name,
		Detail:                  response.Detail,
		CorporationId:           response.CorporationId,
		DefaultMaleRoutineIds:   response.DefaultMaleRoutineIds,
		MaleRoutineIds:          response.MaleRoutineIds,
		DefaultFemaleRoutineIds: response.DefaultFemaleRoutineIds,
		FemaleRoutineIds:        response.FemaleRoutineIds,
		DefaultMaleAvatarId:     response.DefaultMaleAvatarId,
		DefaultFemaleAvatarId:   response.DefaultFemaleAvatarId,
		AvatarIds:               response.AvatarIds,
		ServiceTemplateCategory: response.ServiceTemplateCategory,
		Backgrounds:             backgroundName,
		CreatedBy:               response.CreatedBy,
		DefaultBackground:       fmt.Sprintf("%s/%s/%s", cfg.S3Uri, config.BackgroundBucketName, response.DefaultBackground),
		AvatarUri:               "http://13.113.115.189:4566/routine/defaultRoutineImage.png",
		CreatedAt:               response.CreatedAt,
		UpdatedAt:               response.UpdatedAt,
	}
	return output, nil
}
