package upload

import (
	"avatar/pkg/util"
	"avatar/services/upload/infra/s3/repository"
	"mime/multipart"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
)

type UploadInput struct {
	File   *multipart.FileHeader `form:"file" binding:"required"`
	Bucket string
}

func UploadToS3(input *UploadInput, s3Session *session.Session) error {
	repository := repository.NewS3Repository(s3Session)
	data, err := util.ReadFile(input.File)
	if err != nil {
		return err
	}
	file := strings.NewReader(string(data))
	err = repository.Upload(file, input.File.Filename, input.Bucket)
	if err != nil {
		return err
	}
	return nil
}
