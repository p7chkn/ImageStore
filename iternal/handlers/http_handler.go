package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
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
	pathToFile    string
}

func HttpHandlerNew(repo RepositoryInterface, log *zap.SugaredLogger, ServerAddress string, pathToFile string) *Handler {
	return &Handler{
		repo:          repo,
		log:           log,
		ServerAddress: ServerAddress,
		pathToFile:    pathToFile,
	}
}

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func (h *Handler) GetImage(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "imageName")
	fileBytes, err := ioutil.ReadFile(h.pathToFile + name)
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
