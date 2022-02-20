package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"goImageStore/iternal/utils"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type ErrorMessage struct {
	Details string `json:"details"`
}

type SaveResponse struct {
	Link string `json:"link"`
}

type Handler struct {
	repo          RepositoryInterface
	log           *zap.SugaredLogger
	ServerAddress string
}

func HttpHandlerNew(repo RepositoryInterface, log *zap.SugaredLogger, ServerAddress string) *Handler {
	return &Handler{
		repo:          repo,
		log:           log,
		ServerAddress: ServerAddress,
	}
}

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func (h *Handler) SaveImage(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		h.handleError(w, err)
		return
	}

	file, headers, err := r.FormFile("file")
	if err != nil {
		h.handleError(w, err)
		return
	}
	defer file.Close()

	fileName, err := utils.FormatFileName(headers.Filename)

	if err != nil {
		h.handleError(w, err)
		return
	}

	f, err := os.OpenFile("./downloaded/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		h.log.Error(err)
		return
	}
	defer f.Close()
	if _, err := io.Copy(f, file); err != nil {
		h.handleError(w, err)
		return
	}
	message := SaveResponse{
		Link: "http://" + h.ServerAddress + "/api/image/" + fileName,
	}
	h.handleResponse(w, message)
}

func (h *Handler) GetImage(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "imageName")
	fileBytes, err := ioutil.ReadFile(fmt.Sprintf("./downloaded/%v", name))
	if err != nil {
		h.log.Error(err)
		return
	}
	w.Header().Set("Content-Type", "image/png")
	w.Write(fileBytes)

}

func (h *Handler) handleError(w http.ResponseWriter, err error) {
	message := ErrorMessage{
		Details: err.Error(),
	}
	h.log.Error(err)
	w.WriteHeader(http.StatusBadRequest)
	h.handleResponse(w, message)
}

func (h *Handler) handleResponse(w http.ResponseWriter, message interface{}) {
	response, _ := json.Marshal(message)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
