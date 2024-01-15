package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"path/filepath"
	"slices"

	"github.com/sebasvil20/juansetech-files/src/services"
)

type IUploaderController interface {
	UploadFile(w http.ResponseWriter, r *http.Request)
}

type UploaderController struct {
	UploaderSrv services.IUploaderService
}

func (ctrl *UploaderController) UploadFile(w http.ResponseWriter, r *http.Request) {
	_, fileHeader, err := r.FormFile("file")
	if err != nil {
		handleResponse(w, nil, err, http.StatusBadRequest)
		return
	}

	if !isImage(fileHeader.Filename) {
		handleResponse(w, nil, errors.New("unsupported file type"), http.StatusBadRequest)
		return
	}

	fileURL, err := ctrl.UploaderSrv.UploadFile(r.Context(), fileHeader)
	if err != nil {
		handleResponse(w, nil, err, http.StatusInternalServerError)
		return
	}

	handleResponse(w, fileURL, nil, http.StatusOK)
}

func isImage(fileName string) bool {
	supportedImgsExts := []string{
		".png",
		".jpg",
		".jpeg",
		".webp",
	}
	ext := filepath.Ext(fileName)
	return slices.Contains(supportedImgsExts, ext)
}

func handleResponse(w http.ResponseWriter, data any, err error, statusCode int) {
	type resp struct {
		Data  any    `json:"data,omitempty"`
		Error string `json:"error,omitempty"`
	}
	response := resp{
		Data: data,
	}

	if err != nil {
		response.Error = err.Error()
	}

	marshaledResp, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(marshaledResp)
}
