package avatar

import (
	"avatar/services/gateway/config"
	upload "avatar/services/gateway/infra/upload/respository"
	pb "avatar/services/gateway/protos/center"
	"context"
	"fmt"
	"mime/multipart"
)

type UpdateAvatarInput struct {
	Id        int64                 `json:"id"`
	Name      string                `form:"name"`
	Detail    string                `form:"detail"`
	Image     *multipart.FileHeader `form:"image"`
	Vrm       *multipart.FileHeader `form:"vrm"`
	StartDate string                `form:"startDate"`
	EndDate   string                `form:"endDate"`
	Version   string                `form:"version"`
	Exporter  string                `form:"exporter"`
	Gender    int64                 `form:"gender" binding:"min=0,max=1"`
}

func UpdateAvatar(
	input *UpdateAvatarInput,
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
	var imageLink string
	var vrm string
	if input.Image != nil {
		imageLink = input.Image.Filename
	}
	if input.Vrm != nil {
		vrm = input.Vrm.Filename
	}
	data := &pb.CreateAvatarRequest{
		Id:        input.Id,
		Name:      input.Name,
		Detail:    input.Detail,
		Image:     imageLink,
		Vrm:       vrm,
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
		Gender:    input.Gender,
		Version:   input.Version,
		Exporter:  input.Exporter,
	}
	response, err := centerClient.UpdateAvatar(context.Background(), data)
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
