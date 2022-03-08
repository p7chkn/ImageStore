package services

import (
	"goImageStore/iternal/handlers"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func SetUpRouter(repo handlers.RepositoryInterfaceHttp, log *zap.SugaredLogger, ServerAddress string, pathToFile string) *chi.Mux {
	r := chi.NewRouter()
	handler := handlers.HttpHandlerNew(repo, log, ServerAddress, pathToFile)
	r.Get("/file/ping", handler.Ping)
	r.Get("/file/image/{imageName}", handler.GetImage)
	return r
}
