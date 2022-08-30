package avatar

import (
	"avatar/services/gateway/config"
	upload "avatar/services/gateway/infra/upload/respository"
	pb "avatar/services/gateway/protos/center"
	"context"
	"fmt"
	"mime/multipart"
)

type CreateAvatarInput struct {
	Name      string                `form:"name" binding:"required"`
	Detail    string                `form:"detail" binding:"required"`
	Image     *multipart.FileHeader `form:"image" binding:"required"`
	Vrm       *multipart.FileHeader `form:"vrm" binding:"required"`
	StartDate string                `form:"startDate" binding:"required"`
	EndDate   string                `form:"endDate" binding:"required"`
	Version   string                `form:"version" binding:"required"`
	Exporter  string                `form:"exporter" binding:"required"`
	Gender    int64                 `form:"gender" binding:"min=0,max=1"`
}

type CreateAvatarOutPut struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Detail    string `json:"detail"`
	Image     string `json:"image"`
	Vrm       string `json:"vrm"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Gender    int64  `json:"gender"`
	Version   string `json:"version"`
	Exporter  string `json:"exporter"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func CreateAvatar(
	input *CreateAvatarInput,
	centerClient pb.CenterClient,
	cfg *config.Environment,
	upload upload.UploadClient,
) (*CreateAvatarOutPut, error) {
	if input.Vrm != nil {
		vrm, err := upload.UploadToS3(input.Vrm, config.AvatarBucketName)
		if err != nil {
			return nil, err
		}
		input.Vrm.Filename = vrm
	}
	if input.Image != nil {
		image, err := upload.UploadToS3(input.Image, config.AvatarBucketName)
		if err != nil {
			return nil, err
		}
		input.Image.Filename = image
	}
	data := &pb.CreateAvatarRequest{
		Name:      input.Name,
		Detail:    input.Detail,
		Image:     input.Image.Filename,
		Vrm:       input.Vrm.Filename,
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
		Gender:    input.Gender,
		Version:   input.Version,
		Exporter:  input.Exporter,
	}
	response, err := centerClient.CreateAvatar(context.Background(), data)
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
	}
	return output, nil
}
