package util

import (
	"io/ioutil"
	"mime/multipart"
)

func ReadFile(file *multipart.FileHeader) ([]byte, error) {
	key, err := file.Open()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(key)
	if err != nil {
		return nil, err
	}
	return data, nil
}
