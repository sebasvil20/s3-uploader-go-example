package app

import (
	"log"
	"net/http"

	"github.com/sebasvil20/juansetech-files/src/app/providers"
)

func Run() {
	s3Client, err := providers.ProvideS3Client()
	if err != nil {
		panic(err)
	}

	uploaderSrv := providers.ProvideUploaderService(s3Client)
	uploaderController := providers.ProvideUploaderController(uploaderSrv)

	r := providers.ProvideRouter()
	providers.RegisterRoutes(r, uploaderController)

	log.Print("Server running on port 3000")
	http.ListenAndServe(":3000", r)
}
