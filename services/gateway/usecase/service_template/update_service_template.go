package service_template

import (
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/center"
	"context"
	"fmt"
	"strings"
)

type UpdateServiceTemplateInput struct {
	Id                      int64    `json:"id"`
	Name                    string   `form:"name"`
	Detail                  string   `form:"detail"`
	CorporationId           int64    `form:"corporationId"`
	DefaultMaleRoutineIds   []int64  `form:"defaultMaleRoutineIds"`
	MaleRoutineIds          []int64  `form:"maleRoutineIds"`
	DefaultFemaleRoutineIds []int64  `form:"defaultFemaleRoutineIds"`
	FemaleRoutineIds        []int64  `form:"femaleRoutineIds"`
	DefaultMaleAvatarId     int64    `form:"defaultMaleAvatarId"`
	DefaultFemaleAvatarId   int64    `form:"defaultFemaleAvatarId"`
	AvatarIds               []int64  `form:"avatarIds"`
	ServiceTemplateCategory int64    `form:"serviceTemplateCategory"`
	Backgrounds             []string `form:"backgrounds"`
	DefaultBackground       string   `form:"defaultBackground"`
	AvatarUri               string   `json:"avatarUri"`
}

func UpdateServiceTemplate(
	input *UpdateServiceTemplateInput,
	centerClient pb.CenterClient,
	cfg *config.Environment,
) (*CreateServiceTemplateOutput, error) {
	ctx := context.Background()
	var background []string
	for _, val := range input.Backgrounds {
		splBackground := strings.LastIndex(val, "/")
		background = append(background, val[splBackground+1:])
	}
	in := &pb.UpdateServiceTemplateRequest{
		Id:                      input.Id,
		Name:                    input.Name,
		Detail:                  input.Detail,
		CorporationId:           input.CorporationId,
		MaleRoutineIds:          input.MaleRoutineIds,
		DefaultMaleRoutineIds:   input.DefaultMaleRoutineIds,
		FemaleRoutineIds:        input.FemaleRoutineIds,
		DefaultFemaleRoutineIds: input.DefaultFemaleRoutineIds,
		DefaultMaleAvatarId:     input.DefaultMaleAvatarId,
		DefaultFemaleAvatarId:   input.DefaultFemaleAvatarId,
		AvatarIds:               input.AvatarIds,
		CreatedBy:               0,
		ServiceTemplateCategory: input.ServiceTemplateCategory,
		Backgrounds:             background,
		DefaultBackground:       input.DefaultBackground[strings.LastIndex(input.DefaultBackground, "/")+1:],
	}
	response, err := centerClient.UpdateServiceTemplate(ctx, in)
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
