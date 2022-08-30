package repository

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func NewS3Repository(s3Session *session.Session) *s3RepositoryImpl {
	return &s3RepositoryImpl{
		s3Session: s3Session,
	}
}

type s3RepositoryImpl struct {
	s3Session *session.Session
}

func (r *s3RepositoryImpl) Upload(file *strings.Reader, filename string, bucket string) error {
	uploader := s3manager.NewUploader(r.s3Session)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		return err
	}
	return nil
}
