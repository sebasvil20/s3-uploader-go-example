package services

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"path"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/sebasvil20/juansetech-files/src/config"
)

type IUploaderService interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader) (string, error)
}

type UploaderService struct {
	S3Client *s3.Client
}

func (srv *UploaderService) UploadFile(ctx context.Context, file *multipart.FileHeader) (string, error) {
	openedFile, _ := file.Open()
	fileContent, _ := io.ReadAll(openedFile)
	fileName := srv.generateFileName()
	fileNameWithExt := fmt.Sprintf("%s%s", fileName, path.Ext(file.Filename))

	object := s3.PutObjectInput{
		Bucket:      aws.String(config.BucketS3Name),
		Key:         aws.String(fileNameWithExt),
		Body:        bytes.NewReader(fileContent),
		ContentType: aws.String(fmt.Sprintf("image/%v", path.Ext(file.Filename)[1:])),
	}
	_, err := srv.S3Client.PutObject(ctx, &object)
	if err != nil {
		return "", fmt.Errorf("couldn't upload file, try again later or contact admins - %s", err.Error())
	}

	return fmt.Sprintf("%s/%s", config.CDNURLPrefix, fileNameWithExt), nil
}

func (srv *UploaderService) generateFileName() string {
	randId, _ := gonanoid.Generate("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz", 22)
	return "f_" + randId
}
