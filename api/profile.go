package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func (api *API) ImgProfileView(w http.ResponseWriter, r *http.Request) {
	fileBytes, err := ioutil.ReadFile("./assets/images/img-avatar.png")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return
}

func (api *API) ImgProfileUpdate(w http.ResponseWriter, r *http.Request) {
	// Update image `img-avatar.png` from path `assets/images`
	if err := r.ParseMultipartForm(1024); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	uploadedFile, _, err := r.FormFile("file-avatar")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	//mengambil relative path dari proyek
	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//membentuk lokasi tempat menyimpan file
	fileLocation := filepath.Join(dir, "assets", "images", "img-avatar.png")
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	//mengisi file baru dengan data dari file yang ter-upload
	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	api.dashboardView(w, r)
}
