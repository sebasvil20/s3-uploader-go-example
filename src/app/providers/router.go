package providers

import (
	"github.com/go-chi/chi/v5"
	"github.com/sebasvil20/juansetech-files/src/controllers"
	"net/http"
)

func ProvideRouter() *chi.Mux {
	return chi.NewRouter()
}

func RegisterRoutes(r *chi.Mux, uploaderController controllers.IUploaderController) {
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Post("/upload", uploaderController.UploadFile)
}
