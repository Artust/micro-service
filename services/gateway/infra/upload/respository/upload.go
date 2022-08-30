package repository

import (
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/go-resty/resty/v2"
)

func NewUploadRepository(uri, bucket string) *UploadClient {
	return &UploadClient{
		uri: uri,
	}
}

type UploadClient struct {
	uri string
}

type Response struct {
	Url string
}

func (u *UploadClient) UploadToS3(file *multipart.FileHeader, bucket string) (string, error) {
	newFile, _ := file.Open()
	client := resty.New()
	var response Response
	resp, err := client.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetFileReader("file", file.Filename, newFile).
		SetResult(&response).
		Post(fmt.Sprintf("http://%s/api/s3/%s", u.uri, bucket))
	if err != nil {
		return "", err
	}
	if resp.StatusCode() != 200 {
		return "", errors.New(string(resp.Body()))
	}
	return response.Url, nil
}
