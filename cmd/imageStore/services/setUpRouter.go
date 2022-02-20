package services

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"goImageStore/iternal/handlers"
)

func SetUpRouter(repo handlers.RepositoryInterface, log *zap.SugaredLogger, ServerAddress string) *chi.Mux {
	r := chi.NewRouter()
	handler := handlers.HttpHandlerNew(repo, log, ServerAddress)
	r.Get("/api/ping", handler.Ping)
	r.Post("/api/image/save", handler.SaveImage)
	r.Get("/api/image/{imageName}", handler.GetImage)
	return r
}
