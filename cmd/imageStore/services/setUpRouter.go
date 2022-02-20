package services

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"goImageStore/iternal/handlers"
)

func SetUpRouter(repo handlers.RepositoryInterfaceHttp, log *zap.SugaredLogger, ServerAddress string, pathToFile string) *chi.Mux {
	r := chi.NewRouter()
	handler := handlers.HttpHandlerNew(repo, log, ServerAddress, pathToFile)
	r.Get("/file/ping", handler.Ping)
	r.Get("/file/image/{imageName}", handler.GetImage)
	return r
}
