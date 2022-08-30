package service_template

import (
	"avatar/services/gateway/config"
	upload "avatar/services/gateway/infra/upload/respository"
	"avatar/services/gateway/pkg/util"
	pb "avatar/services/gateway/protos/center"
	"context"
	"fmt"
	"mime/multipart"
	"strings"

	"golang.org/x/sync/errgroup"
)

type CreateServiceTemplateInput struct {
	Name                    string                  `form:"name" binding:"required"`
	Detail                  string                  `form:"detail" binding:"required"`
	CorporationId           int64                   `form:"corporationId" binding:"required"`
	DefaultMaleRoutineIds   []int64                 `form:"defaultMaleRoutineIds" binding:"required"`
	MaleRoutineIds          []int64                 `form:"maleRoutineIds" binding:"required"`
	DefaultFemaleRoutineIds []int64                 `form:"defaultFemaleRoutineIds" binding:"required"`
	FemaleRoutineIds        []int64                 `form:"femaleRoutineIds" binding:"required"`
	DefaultMaleAvatarId     int64                   `form:"defaultMaleAvatarId" binding:"required"`
	DefaultFemaleAvatarId   int64                   `form:"defaultFemaleAvatarId" binding:"required"`
	AvatarIds               []int64                 `form:"avatarIds" binding:"required"`
	ServiceTemplateCategory int64                   `form:"serviceTemplateCategory" binding:"required"`
	Backgrounds             []*multipart.FileHeader `form:"backgrounds" binding:"required"`
	DefaultBackground       string                  `form:"defaultBackground" binding:"required"`
}

type CreateServiceTemplateOutput struct {
	Id                      int64    `json:"id"`
	Name                    string   `form:"name"`
	Detail                  string   `form:"detail"`
	CorporationId           int64    `form:"corporationId"`
	MaleRoutineIds          []int64  `form:"maleRoutineIds"`
	DefaultMaleRoutineIds   []int64  `form:"defaultMaleRoutineIds"`
	FemaleRoutineIds        []int64  `form:"femaleRoutineIds"`
	DefaultFemaleRoutineIds []int64  `form:"defaultFemaleRoutineIds"`
	AvatarIds               []int64  `form:"avatarIds"`
	DefaultMaleAvatarId     int64    `form:"defaultMaleAvatarId"`
	DefaultFemaleAvatarId   int64    `form:"defaultFemaleAvatarId"`
	ServiceTemplateCategory int64    `form:"serviceTemplateCategory"`
	Backgrounds             []string `form:"backgrounds"`
	DefaultBackground       string   `form:"defaultBackground"`
	AvatarUri               string   `json:"avatarUri"`
	CreatedBy               int64    `json:"createdBy"`
	CreatedAt               string   `json:"createdAt"`
	UpdatedAt               string   `json:"updatedAt,omitempty"`
}

func CreateServiceTemplate(
	input *CreateServiceTemplateInput,
	centerClient pb.CenterClient,
	upload upload.UploadClient,
	cfg *config.Environment,
) (*CreateServiceTemplateOutput, error) {
	ctx, cancel := context.WithCancel(context.Background())
	eg, _ := errgroup.WithContext(ctx)
	var err error
	file := make(chan *multipart.FileHeader)
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case data := <-file:
				_, err = upload.UploadToS3(data, config.BackgroundBucketName)
				if err != nil {
					return err
				}
			}
		}
	})
	var backgroundName []string
	for _, background := range input.Backgrounds {
		fileName, _ := util.GenerateFileName(background.Filename, "")
		fileName = strings.Trim(fileName, "-")
		if background.Filename == input.DefaultBackground {
			input.DefaultBackground = fileName
		}
		background.Filename = fileName
		backgroundName = append(backgroundName, fileName)
		file <- background
	}
	cancel()
	if eg.Wait() != nil {
		return nil, eg.Wait()
	}
	createServiceTemplateInput := &pb.CreateServiceTemplateRequest{
		Name:                    input.Name,
		Detail:                  input.Detail,
		CorporationId:           input.CorporationId,
		MaleRoutineIds:          input.MaleRoutineIds,
		DefaultMaleRoutineIds:   input.FemaleRoutineIds,
		FemaleRoutineIds:        input.FemaleRoutineIds,
		DefaultFemaleRoutineIds: input.DefaultFemaleRoutineIds,
		DefaultMaleAvatarId:     input.DefaultMaleAvatarId,
		DefaultFemaleAvatarId:   input.DefaultFemaleAvatarId,
		AvatarIds:               input.AvatarIds,
		CreatedBy:               1,
		ServiceTemplateCategory: input.ServiceTemplateCategory,
		Backgrounds:             backgroundName,
		DefaultBackground:       input.DefaultBackground,
	}
	response, err := centerClient.CreateServiceTemplate(context.Background(), createServiceTemplateInput)
	if err != nil {
		return nil, err
	}
	backgroundName = []string{}
	for _, background := range response.Backgrounds {
		backgroundName = append(backgroundName, fmt.Sprintf("%s/%s/%s", cfg.S3Uri, config.BackgroundBucketName, background))
	}
	createServiceTemplateOutput := &CreateServiceTemplateOutput{
		Id:                      response.Id,
		Name:                    response.Name,
		Detail:                  response.Detail,
		CorporationId:           response.CorporationId,
		MaleRoutineIds:          response.MaleRoutineIds,
		DefaultMaleRoutineIds:   response.DefaultMaleRoutineIds,
		FemaleRoutineIds:        response.FemaleRoutineIds,
		DefaultFemaleRoutineIds: response.DefaultFemaleRoutineIds,
		AvatarIds:               response.AvatarIds,
		DefaultMaleAvatarId:     response.DefaultMaleAvatarId,
		DefaultFemaleAvatarId:   response.DefaultFemaleAvatarId,
		ServiceTemplateCategory: response.ServiceTemplateCategory,
		Backgrounds:             backgroundName,
		DefaultBackground:       fmt.Sprintf("%s/%s/%s", cfg.S3Uri, config.BackgroundBucketName, response.DefaultBackground),
		AvatarUri:               "http://13.113.115.189:4566/routine/defaultRoutineImage.png",
		CreatedBy:               response.CreatedBy,
		CreatedAt:               response.CreatedAt,
		UpdatedAt:               "",
	}
	return createServiceTemplateOutput, nil
}
