package upload_file

import (
	upload "avatar/services/gateway/infra/upload/respository"
	"avatar/services/gateway/pkg/util"
	"mime/multipart"
	"strings"
)

type UploadFileInput struct {
	Bucket     string                `form:"bucket"`
	Background *multipart.FileHeader `form:"file" binding:"required"`
}

type UploadFileOutput struct {
	Name string `json:"name"`
}

func UploadFile(
	input *UploadFileInput,
	upload upload.UploadClient,
) (*UploadFileOutput, error) {
	fileName, _ := util.GenerateFileName(input.Background.Filename, "")
	input.Background.Filename = strings.Trim(fileName, "-")
	fileName, err := upload.UploadToS3(input.Background, input.Bucket)
	if err != nil {
		return nil, err
	}
	output := &UploadFileOutput{
		Name: fileName,
	}
	return output, nil
}
