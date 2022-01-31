package services

import (
	"gin-cloudinary-api/helper"
	"gin-cloudinary-api/models"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

type media interface {
	FileUpload(file models.File) (string, error)
	RemoteUpload(url models.Url) (string, error)
}

type mediaUpload struct {}

func NewMediaUpload() media {
	return &mediaUpload{}
}

func (*mediaUpload) FileUpload(file models.File) (string, error) {
	//validate
	err := validate.Struct(file)
	if err != nil {
		return "", err
	}
	
	//upload
	uploadUrl, err := helper.ImageUploadHelper(file.File)
	if err != nil {
		return "", err
	}

	return uploadUrl, nil
}

func (*mediaUpload) RemoteUpload(url models.Url) (string, error) {
	//validate
	err := validate.Struct(url)
	if err != nil {
		return "", err
	}
	
	//upload
	uploadUrl, errUrl := helper.ImageUploadHelper(url.Url)
	if errUrl != nil {
		return "", err
	}

	return uploadUrl, nil
}