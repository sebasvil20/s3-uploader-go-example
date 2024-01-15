package providers

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sebasvil20/juansetech-files/src/controllers"
	"github.com/sebasvil20/juansetech-files/src/services"
)

func ProvideUploaderService(s3Client *s3.Client) services.IUploaderService {
	return &services.UploaderService{
		S3Client: s3Client,
	}
}

func ProvideUploaderController(uploaderSrv services.IUploaderService) controllers.IUploaderController {
	return &controllers.UploaderController{
		UploaderSrv: uploaderSrv,
	}
}
